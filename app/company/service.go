package company

type Service struct {
	factory FactoryInterface
}

type ServiceInterface interface {
	CreateCompany(name, addressLine1, addressLine2 string) Company
}

func NewService(factory FactoryInterface) *Service {
	return &Service{factory: factory}
}

func (s *Service) CreateCompany(name, addressLine1, addressLine2 string) Company {
	return s.factory.Create(name, addressLine1, addressLine2)
}
