package auth

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/user"
	"net/http"
	"time"
)

type Service struct {
	userService     user.ServiceInterface
	googleProvider  *google.Provider
	mgoogleProvider *google.Provider
	OauthServer     OauthServerInterface
}

type ServiceInterface interface {
	HandleLogin(ctx *gin.Context)
	GetAuthLoginUrl(ctx *gin.Context) (string, core.DetailedError)
	//Authorize(ctx *gin.Context) (u user.User, err core.DetailedError)
	HandleLogout(ctx *gin.Context) core.DetailedError
	HandleCallback(ctx *gin.Context) (user.User, core.DetailedError)
}

func NewService(userService user.ServiceInterface, oauthServer OauthServerInterface) *Service {
	provider := google.New(
		core.EnvConfigs.GoogleOauthClientID,
		core.EnvConfigs.GoogleOauthClientSecret,
		core.EnvConfigs.GoogleOauthRedirectUri,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	)
	mobileProvider := google.New(
		core.EnvConfigs.GoogleOauthClientID,
		core.EnvConfigs.GoogleOauthClientSecret,
		core.EnvConfigs.MobileGoogleOauthRedirectUri,
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	)
	mobileProvider.SetName("mgoogle")
	goth.UseProviders(
		provider,
		mobileProvider,
	)

	oauthServer.SaveClient(provider.Name(), provider.Secret, provider.CallbackURL)
	oauthServer.SaveClient(mobileProvider.Name(), mobileProvider.Secret, mobileProvider.CallbackURL)

	return &Service{userService, provider, mobileProvider, oauthServer}
}

func (s *Service) HandleLogin(ctx *gin.Context) {
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (s *Service) GetAuthLoginUrl(ctx *gin.Context) (string, core.DetailedError) {
	url, err := gothic.GetAuthURL(ctx.Writer, ctx.Request)
	if err != nil {
		return url, core.NewBadRequestError(err.Error())
	}
	return url, nil
}

func (s *Service) HandleLogout(ctx *gin.Context) core.DetailedError {
	err := gothic.Logout(ctx.Writer, ctx.Request)
	sessions.Default(ctx).Clear()
	if err != nil {
		return core.NewSimpleError(err.Error())
	}
	return nil
}

//func (s *Service) Authorize(ctx *gin.Context) (u user.User, err core.DetailedError) {
//	return s.HandleCallback(ctx)
//}

func (s *Service) HandleCallback(ctx *gin.Context) (u user.User, err core.DetailedError) {

	provider, pErr := gothic.GetProviderName(ctx.Request)
	if pErr != nil {
		return u, core.NewUnauthorizedError(pErr.Error())
	}

	authUser, authErr := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if authErr != nil {
		return u, core.NewUnauthorizedError(authErr.Error())
	}

	err = s.saveTokenData(provider, authUser)
	if err != nil {
		return u, err
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
	err = s.addUserToSession(ctx, authUser)
	if err != nil {
		return u, err
	}

	if provider == "mgoogle" {
		// Redirect to the mobile app with the authorization code
		mobileRedirectURI := fmt.Sprintf("com.gshot.garage:/auth/mgoogle/callback?userId=%s&token=%s", u.ID, authUser.AccessToken)
		ctx.Redirect(http.StatusSeeOther, mobileRedirectURI)
		return
	}
	return u, nil
}

func (s *Service) saveTokenData(provider string, authUser goth.User) core.DetailedError {
	p, pErr := goth.GetProvider(provider)
	if pErr != nil {
		return core.NewSimpleError(pErr.Error())
	}

	tokenData := osin.AccessData{
		Client: &osin.DefaultClient{
			Id: p.Name(),
		},
		AccessToken:  authUser.AccessToken,
		RefreshToken: authUser.RefreshToken,
		ExpiresIn:    int32(authUser.ExpiresAt.Sub(time.Now()).Seconds()),
		Scope:        "https://www.googleapis.com/auth/userinfo.email,https://www.googleapis.com/auth/userinfo.profile",
		RedirectUri:  s.googleProvider.CallbackURL,
		CreatedAt:    time.Now(),
		UserData: map[string]string{
			"user_id": authUser.UserID,
		},
	}
	err := s.OauthServer.SaveToken(&tokenData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) addUserToSession(ctx *gin.Context, user goth.User) core.DetailedError {
	session := sessions.Default(ctx)

	// Remove the raw data to reduce the size
	user.RawData = map[string]interface{}{}

	session.Set("user", user)
	err := session.Save()

	if err != nil {
		return core.NewSimpleError(err.Error())
	}
	return nil
}
