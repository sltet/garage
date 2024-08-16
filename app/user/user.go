package user

import (
	"github.com/sltet/garage/app/core"
)

type User struct {
	core.Entity
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
} //@name User

func NewUser(firstName string, lastName string, email string, externalId string) User {
	return User{
		core.NewEntity(),
		firstName,
		lastName,
		email,
		externalId,
	}
}
