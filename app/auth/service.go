package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/user"
	"log"
)

type Service struct {
	userService    user.ServiceInterface
	googleProvider *google.Provider
}

type ServiceInterface interface {
	HandleLogin(ctx *gin.Context)
	HandleLogout(ctx *gin.Context) core.DetailedError
	HandleCallback(ctx *gin.Context) (user.User, core.DetailedError)
}

func NewService(userService user.ServiceInterface) *Service {
	provider := google.New(
		core.EnvConfigs.GoogleOauthClientID,
		core.EnvConfigs.GoogleOauthClientSecret,
		core.EnvConfigs.GoogleOauthRedirectUri,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	)
	goth.UseProviders(
		provider,
	)

	return &Service{userService, provider}
}

func (s *Service) HandleLogin(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (s *Service) HandleLogout(ctx *gin.Context) core.DetailedError {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	err := gothic.Logout(ctx.Writer, ctx.Request)
	if err != nil {
		return core.NewSimpleError(err.Error())
	}
	return nil
}

func (s *Service) HandleCallback(ctx *gin.Context) (u user.User, err core.DetailedError) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	authUser, authErr := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if authErr != nil {
		return u, core.NewUnauthorizedError(authErr.Error())
	}

	userToCreate := user.UserCreate{
		FirstName:  authUser.Name,
		LastName:   authUser.LastName,
		Email:      authUser.Email,
		ExternalId: authUser.UserID,
	}

	u, err = s.userService.CreateUser(ctx, userToCreate)
	if err != nil {
		return u, err
	}
	// Store user data in the session
	addUserToSession(ctx, authUser)
	return u, nil
}

func addUserToSession(ctx *gin.Context, user goth.User) {
	session := sessions.Default(ctx)

	// Remove the raw data to reduce the size
	user.RawData = map[string]interface{}{}

	session.Set("user", user)
	session.Options(sessions.Options{
		MaxAge: 3600 * 12, // 12hrs
	})
	err := session.Save()

	if err != nil {
		log.Print("Problem Saving session data", err)
	}
}
