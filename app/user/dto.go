package user

type UserCreate struct {
	FirstName string `json:"first_name" binding:"required,min=5,max=20"`
	LastName  string `json:"last_name" binding:"required"`
}
