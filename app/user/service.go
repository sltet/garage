package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
)

type Service struct {
	factory    FactoryInterface
	repository RepositoryInterface
}

type ServiceInterface interface {
	CreateUser(ctx *gin.Context, u UserCreate) (user User, err core.DetailedError)
	UpdateUser(ctx *gin.Context, id string, u UserUpdate) (user User, err core.DetailedError)
	FindAll(ctx *gin.Context) ([]User, core.DetailedError)
	FindById(ctx *gin.Context, id string) (User, core.DetailedError)
	FindByExternalId(ctx *gin.Context, id string) (User, core.DetailedError)
}

func NewService(factory FactoryInterface, repository RepositoryInterface) *Service {
	return &Service{factory: factory, repository: repository}
}

func (s Service) FindAll(ctx *gin.Context) ([]User, core.DetailedError) {
	return s.repository.FindAll(ctx)
}

func (s Service) FindById(ctx *gin.Context, id string) (User, core.DetailedError) {
	return s.repository.FindById(ctx, id)
}

func (s Service) FindByExternalId(ctx *gin.Context, id string) (User, core.DetailedError) {
	return s.repository.FindById(ctx, id)
}

func (s Service) CreateUser(ctx *gin.Context, u UserCreate) (user User, err core.DetailedError) {
	user = s.factory.CreateUser(ctx, u)
	user, err = s.repository.Create(ctx, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s Service) UpdateUser(ctx *gin.Context, id string, u UserUpdate) (user User, err core.DetailedError) {
	user, err = s.repository.FindById(ctx, id)
	if err != nil {
		return user, err
	}
	updatedUser := s.factory.UpdateUser(ctx, user, u)
	return s.repository.Save(ctx, updatedUser)
}
