package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
)

type Service struct {
	factory FactoryInterface
}

type ServiceInterface interface {
	CreateUser(ctx *gin.Context) (User, error)
	FindAll(ctx *gin.Context) []User
}

func NewService(factory FactoryInterface) *Service {
	return &Service{factory: factory}
}

func (s *Service) FindAll(ctx *gin.Context) []User {
	return []User{
		NewUser("steve", "landry"),
	}
}

func (s *Service) CreateUser(ctx *gin.Context) (user User, err error) {
	user, err = s.factory.CreateUser(ctx)
	if err != nil {
		return user, err
	}
	user.ID = core.GetTimeBasedUUID().String()
	return user, err
}
