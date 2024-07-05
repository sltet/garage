package appointment

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

type FactoryInterface interface {
	Create(name, addressLine1, addressLine2 string) Appointment
}

func (f Factory) Create(customerID, locationID, serviceID string) Appointment {
	return NewAppointment(customerID, locationID, serviceID)
}
