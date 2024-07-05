package appointment

import (
	"github.com/sltet/garage/app/core"
)

type Appointment struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	LocationID string `json:"location_id"`
	ServiceID  string `json:"service_id"`
}

func NewAppointment(customerID, locationID, serviceID string) Appointment {
	return Appointment{
		ID:         core.GetTimeBasedUUID().String(),
		CustomerID: customerID,
		LocationID: locationID,
		ServiceID:  serviceID,
	}
}
