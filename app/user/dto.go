package user

// UserCreate model info
//
//	@Description	UserCreate information
//	@Description	with user firstname and lastname
type UserCreate struct {
	FirstName  string `json:"first_name" example:"john" binding:"required,min=2,max=20"`
	LastName   string `json:"last_name" example:"doe" binding:"required,min=2,max=20"`
	Email      string `json:"email" example:"john@doe.com" binding:"required,email"`
	ExternalId string `json:"external_id" example:"123456" binding:"min=2,max=20"`
} //@name UserCreate

// UserUpdate model info
type UserUpdate struct {
	FirstName string `json:"first_name" example:"john" binding:"required,min=2,max=20"`
	LastName  string `json:"last_name" example:"doe" binding:"required,min=2,max=20"`
	Email     string `json:"email" example:"john@doe.com" binding:"required,email"`
} //@name UserUpdate
