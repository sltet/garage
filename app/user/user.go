package user

import (
	"github.com/sltet/garage/app/core"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUser(firstName string, lastName string, email string) User {
	return User{
		core.GetTimeBasedUUID().String(),
		firstName,
		lastName,
		email,
	}
}

func (c User) GetID() string {
	return c.ID
}
