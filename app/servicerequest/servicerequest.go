package servicerequest

type ServiceRequest struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
