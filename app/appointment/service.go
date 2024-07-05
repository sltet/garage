package appointment

type Service struct {
	factory FactoryInterface
}

type ServiceInterface interface {
	CreateAppointment(customerID, locationID, serviceID string) Appointment
}

func NewService(factory FactoryInterface) *Service {
	return &Service{factory: factory}
}

func (s *Service) CreateAppointment(customerID, locationID, serviceID string) Appointment {
	return s.factory.Create(customerID, locationID, serviceID)
}
