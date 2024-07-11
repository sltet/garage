package user

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	factory    FactoryInterface
	repository RepositoryInterface
}

type ServiceInterface interface {
	CreateUser(ctx *gin.Context, u UserCreate) (user User, err error)
	UpdateUser(ctx *gin.Context, id string, u UserUpdate) (user User, err error)
	FindAll(ctx *gin.Context) ([]User, error)
	FindById(ctx *gin.Context, id string) (User, error)
}

func NewService(factory FactoryInterface, repository RepositoryInterface) *Service {
	return &Service{factory: factory, repository: repository}
}

func (s Service) FindAll(ctx *gin.Context) ([]User, error) {
	return s.repository.FindAll(ctx)
}

func (s Service) FindById(ctx *gin.Context, id string) (User, error) {
	return s.repository.FindById(ctx, id)
}

func (s Service) CreateUser(ctx *gin.Context, u UserCreate) (user User, err error) {
	user = s.factory.CreateUser(ctx, u)
	user, err = s.repository.Create(ctx, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s Service) UpdateUser(ctx *gin.Context, id string, u UserUpdate) (user User, err error) {
	user, err = s.repository.FindById(ctx, id)
	if err != nil {
		return user, err
	}
	updatedUser := s.factory.UpdateUser(ctx, user, u)
	return s.repository.Save(ctx, updatedUser)
}
