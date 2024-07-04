package user

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type ControllerInterface interface {
	FindAllUsers(ctx *gin.Context)
}

// FindAllUsers godoc
// @Summary find all users
// @Schemes
// @Description find all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /users [get]
func (c Controller) FindAllUsers(ctx *gin.Context) {
	ctx.JSON(200, []User{
		NewUser("Steve", "Landry"),
	})
}
