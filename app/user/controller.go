package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{
		service: service,
	}
}

type ControllerInterface interface {
	FindAllUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

// FindAllUsers godoc
// @Summary find all users
// @Schemes
// @Description find all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func (c Controller) FindAllUsers(ctx *gin.Context) {
	users := c.service.FindAll(ctx)
	ctx.JSON(200, users)
}

// CreateUser godoc
// @Summary create user
// @Schemes
// @Description create user
// @Param user body UserCreate true "the user to create"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Failure 400 {object} string
// @Router /users [post]
func (c Controller) CreateUser(ctx *gin.Context) {
	user, err := c.service.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
