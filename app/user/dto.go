package user

// UserCreate model info
// @Description UserCreate information
// @Description with user firstname and lastname
type UserCreate struct {
	// user firstname
	FirstName string `json:"first_name" example:"john" binding:"required,min=2,max=20"`
	// user lastname
	LastName string `json:"last_name" example:"doe" binding:"required,min=2,max=20"`
}

// UserUpdate model info
type UserUpdate struct {
	// user firstname
	FirstName string `json:"first_name" example:"john" binding:"required,min=2,max=20"`
	// user lastname
	LastName string `json:"last_name" example:"doe" binding:"required,min=2,max=20"`
}
