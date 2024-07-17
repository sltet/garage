package company

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
)

type Service struct {
	factory    FactoryInterface
	repository RepositoryInterface
}

type ServiceInterface interface {
	CreateCompany(ctx *gin.Context, company CompanyCreate) (Company, core.DetailedError)
	FindAll(ctx *gin.Context) ([]Company, core.DetailedError)
	FindById(ctx *gin.Context, id string) (Company, core.DetailedError)
}

func NewService(factory FactoryInterface, repository RepositoryInterface) *Service {
	return &Service{factory: factory, repository: repository}
}

func (s *Service) CreateCompany(ctx *gin.Context, company CompanyCreate) (Company, core.DetailedError) {
	comp := s.factory.Create(company)
	return s.repository.Create(ctx, comp)
}

func (s *Service) FindAll(ctx *gin.Context) ([]Company, core.DetailedError) {
	return s.repository.FindAll(ctx)
}

func (s *Service) FindById(ctx *gin.Context, id string) (Company, core.DetailedError) {
	return s.repository.FindById(ctx, id)
}

func (s *Service) Create(ctx *gin.Context) ([]Company, core.DetailedError) {
	return s.repository.FindAll(ctx)
}
