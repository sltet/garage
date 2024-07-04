package user

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUser(firstName string, lastName string) User {
	return User{
		firstName,
		lastName,
	}
}
