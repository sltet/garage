package servicerequest

type ServiceRequest struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (c ServiceRequest) GetID() string {
	return c.ID
}
