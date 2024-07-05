package company

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

type FactoryInterface interface {
	Create(name, addressLine1, addressLine2 string) Company
}

func (f Factory) Create(name, addressLine1, addressLine2 string) Company {
	return NewCompany(name, addressLine1, addressLine2)
}
