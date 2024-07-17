package company

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

type FactoryInterface interface {
	Create(company CompanyCreate) Company
}

func (f Factory) Create(company CompanyCreate) Company {
	return NewCompany(company)
}
