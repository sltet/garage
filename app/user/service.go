package user

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	factory    FactoryInterface
	repository RepositoryInterface
}

type ServiceInterface interface {
	CreateUser(ctx *gin.Context, u UserCreate) User
	UpdateUser(ctx *gin.Context, userId string, u UserUpdate) User
	FindAll(ctx *gin.Context) ([]User, error)
}

func NewService(factory FactoryInterface, repository RepositoryInterface) *Service {
	return &Service{factory: factory, repository: repository}
}

func (s Service) FindAll(ctx *gin.Context) ([]User, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) CreateUser(ctx *gin.Context, u UserCreate) (user User) {
	return s.factory.CreateUser(ctx, u)
}

func (s *Service) UpdateUser(ctx *gin.Context, id string, u UserUpdate) (user User) {
	return s.factory.UpdateUser(ctx, User{ID: id}, u)
}
