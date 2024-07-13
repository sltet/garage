package operation

import "github.com/gin-gonic/gin"

type Service struct {
	repository RepositoryInterface
}

type ServiceInterface interface {
	FindAllOperations(ctx *gin.Context) ([]ServiceOperation, error)
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository}
}

func (s *Service) FindAllOperations(ctx *gin.Context) ([]ServiceOperation, error) {
	return s.repository.FindAll(ctx)
}
