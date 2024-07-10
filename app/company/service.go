package company

import "github.com/gin-gonic/gin"

type Service struct {
	factory    FactoryInterface
	repository RepositoryInterface
}

type ServiceInterface interface {
	CreateCompany(ctx *gin.Context, company CompanyCreate) (Company, error)
	FindAll(ctx *gin.Context) ([]Company, error)
}

func NewService(factory FactoryInterface, repository RepositoryInterface) *Service {
	return &Service{factory: factory, repository: repository}
}

func (s *Service) CreateCompany(ctx *gin.Context, company CompanyCreate) (Company, error) {
	comp := s.factory.Create(company)
	return s.repository.Create(ctx, comp)
}

func (s *Service) FindAll(ctx *gin.Context) ([]Company, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) Create(ctx *gin.Context) ([]Company, error) {
	return s.repository.FindAll(ctx)
}
