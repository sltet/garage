package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	googleauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"net/http"
)

type GoogleService struct {
	authConfig *oauth2.Config
}

type GoogleServiceInterface interface {
	HandleLogin(ctx *gin.Context)
	HandleCallbackLogin(ctx *gin.Context) *googleauth.Userinfo
	//AuthenticateUser(ctx *gin.Context, token string) (*googleauth.Userinfo, core.DetailedError)
	//ValidateToken(ctx *gin.Context, token string) (*googleauth.Tokeninfo, core.DetailedError)
}

func NewGoogleService() *GoogleService {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  core.EnvConfigs.GoogleOauthRedirectUri,
		ClientID:     core.EnvConfigs.GoogleOauthClientID,
		ClientSecret: core.EnvConfigs.GoogleOauthClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return &GoogleService{
		googleOauthConfig,
	}
}

func (s *GoogleService) HandleLogin(ctx *gin.Context) {
	url := s.authConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline, oauth2.ApprovalForce)
	http.Redirect(ctx.Writer, ctx.Request, url, http.StatusTemporaryRedirect)
}

func (s *GoogleService) HandleCallbackLogin(ctx *gin.Context) *googleauth.Userinfo {
	// Read state from Cookie
	oauthState := ctx.Request.FormValue("state")
	if oauthState != oauthStateString {
		fmt.Println("invalid state")
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusTemporaryRedirect)
		return nil
	}

	accessToken, err := s.authConfig.Exchange(ctx, ctx.Request.FormValue("code"))
	if err != nil {
		fmt.Println(fmt.Errorf("code exchange failed: %s", err.Error()))
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusTemporaryRedirect)
		return nil
	}

	userinfo, err := s.getUserInfo(ctx, accessToken)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusTemporaryRedirect)
		return nil
	}

	return userinfo
}

//func (s *GoogleService) AuthenticateUser(ctx *gin.Context, token string) (*googleauth.Userinfo, core.DetailedError) {
//	_, err := s.ValidateToken(ctx, token)
//	if err != nil {
//		return nil, err
//	}
//	userInfo, infoErr := s.userInfoService.Me.Get().Do(googleapi.QueryParameter("id_token", token))
//	if infoErr != nil {
//		e, _ := infoErr.(*googleapi.Error)
//		return nil, core.NewSimpleError(e.Message)
//	}
//	return userInfo, nil
//}
//
//func (s *GoogleService) ValidateToken(ctx *gin.Context, token string) (*googleauth.Tokeninfo, core.DetailedError) {
//	tokenValidator, err := idtoken.NewValidator(ctx)
//	if err != nil {
//		return nil, core.NewSimpleError(err.Error())
//	}
//	_, err = tokenValidator.Validate(ctx, token, core.EnvConfigs.GoogleOauthClientID)
//	if err != nil {
//		return nil, core.NewSimpleError(err.Error())
//	}
//	tokenInfoCall := s.oauth2service.Tokeninfo()
//	tokenInfoCall.IdToken(token)
//	tokenInfo, err := tokenInfoCall.Do()
//	if err != nil {
//		return nil, core.NewBadRequestError(err.Error())
//	}
//	return tokenInfo, nil
//}

//func (s *GoogleService) getUserInfos(token string) (*googleauth.Userinfo, core.DetailedError) {
//	userInfo, infoErr := s.userInfoService..Get().Do(googleapi.QueryParameter("access_token", token))
//	if infoErr != nil {
//		e, _ := infoErr.(*googleapi.Error)
//		return nil, core.NewSimpleError(e.Message)
//	}
//	return userInfo, nil
//}

func (s *GoogleService) getUserInfo(ctx *gin.Context, token *oauth2.Token) (*googleauth.Userinfo, core.DetailedError) {
	userInfoService, err := s.getGoogleUserService(ctx, token)
	if err != nil {
		return nil, err
	}

	userInfo, infoErr := userInfoService.Get().Do(googleapi.QueryParameter("access_token", token.AccessToken))
	if infoErr != nil {
		e, _ := infoErr.(*googleapi.Error)
		return nil, core.NewSimpleError(e.Message)
	}
	return userInfo, nil
}

func (s *GoogleService) getGoogleUserService(ctx *gin.Context, token *oauth2.Token) (*googleauth.UserinfoV2MeService, core.DetailedError) {
	oauth2Service, err := googleauth.NewService(ctx, option.WithTokenSource(s.authConfig.TokenSource(ctx, token)))
	if err != nil {
		return nil, core.NewSimpleError(err.Error())
	}
	return googleauth.NewUserinfoV2MeService(oauth2Service), nil
}
