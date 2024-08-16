package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/user"
)

type Service struct {
	userService   user.ServiceInterface
	googleService GoogleServiceInterface
}

type ServiceInterface interface {
	HandleGoogleLogin(ctx *gin.Context)
	HandleCallbackGoogleLogin(ctx *gin.Context) (user.User, core.DetailedError)
	//AuthenticateGoogleToken(ctx *gin.Context, token string) (u user.User, err core.DetailedError)
}

func NewService(userService user.ServiceInterface, googleService GoogleServiceInterface) *Service {
	return &Service{userService, googleService}
}

func (s *Service) HandleGoogleLogin(ctx *gin.Context) {
	s.googleService.HandleLogin(ctx)
}

func (s *Service) HandleCallbackGoogleLogin(ctx *gin.Context) (user.User, core.DetailedError) {
	userInfo := s.googleService.HandleCallbackLogin(ctx)
	userToCreate := user.UserCreate{
		FirstName:  userInfo.Name,
		LastName:   userInfo.FamilyName,
		Email:      userInfo.Email,
		ExternalId: userInfo.Id,
	}

	return s.userService.CreateUser(ctx, userToCreate)
}

//func (s *Service) AuthenticateGoogleToken(ctx *gin.Context, token string) (u user.User, err core.DetailedError) {
//	userInfo, err := s.googleService.AuthenticateUser(ctx, token)
//	if err != nil {
//		return u, err
//	}
//	if userInfo == nil {
//		return u, core.NewSimpleError("Could not fetch user details from Google")
//	}
//	userToCreate := user.UserCreate{
//		FirstName:  userInfo.Name,
//		LastName:   userInfo.FamilyName,
//		Email:      userInfo.Email,
//		ExternalId: userInfo.Id,
//	}
//
//	return s.userService.CreateUser(ctx, userToCreate)
//}
