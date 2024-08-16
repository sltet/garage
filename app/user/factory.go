package user

import (
	"github.com/gin-gonic/gin"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

type FactoryInterface interface {
	CreateUser(ctx *gin.Context, u UserCreate) (user User)
	UpdateUser(ctx *gin.Context, user User, u UserUpdate) User
}

func (f Factory) convert(user UserCreate) User {
	return NewUser(user.FirstName, user.LastName, user.Email, user.ExternalId)
}

func (f Factory) CreateUser(_ *gin.Context, u UserCreate) (user User) {
	return f.convert(u)
}

func (f Factory) UpdateUser(_ *gin.Context, user User, u UserUpdate) User {
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	return user
}
