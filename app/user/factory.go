package user

import (
	"github.com/gin-gonic/gin"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

type FactoryInterface interface {
	CreateUser(ctx *gin.Context) (user User, err error)
}

func (f Factory) convert(user UserCreate) User {
	return NewUser(user.FirstName, user.LastName)
}

func (f Factory) CreateUser(ctx *gin.Context) (user User, err error) {
	var u UserCreate
	if err := ctx.ShouldBindJSON(&u); err != nil {
		return user, err
	}
	return f.convert(u), nil
}
