package operation

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
)

type Service struct {
	repository RepositoryInterface
}

type ServiceInterface interface {
	FindAllOperations(ctx *gin.Context) ([]ServiceOperation, core.DetailedError)
	FindById(ctx *gin.Context, id string) (ServiceOperation, core.DetailedError)
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository}
}

func (s *Service) FindAllOperations(ctx *gin.Context) ([]ServiceOperation, core.DetailedError) {
	return s.repository.FindAll(ctx)
}

func (s *Service) FindById(ctx *gin.Context, id string) (ServiceOperation, core.DetailedError) {
	return s.repository.FindById(ctx, id)
}
