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
	UpdateUser(ctx *gin.Context)
}

// FindAllUsers godoc
//
//	@Summary	find all users
//	@Schemes
//	@Description	find all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	User
//	@Router			/users [get]
func (c Controller) FindAllUsers(ctx *gin.Context) {
	users, err := c.service.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, users)
}

// CreateUser godoc
//
//	@Summary	create user
//	@Schemes
//	@Description	create user
//	@Param			user	body	UserCreate	true	"the user to create"
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Failure		400	{object}	string
//	@Router			/users [post]
func (c Controller) CreateUser(ctx *gin.Context) {
	var u UserCreate
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user := c.service.CreateUser(ctx, u)
	ctx.JSON(200, user)
}

// UpdateUser godoc
//
//	@Summary	update user
//	@Schemes
//	@Description	update user
//	@Param			id		path	string		true	"the user id"
//	@Param			user	body	UserUpdate	true	"the user to update"
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Failure		400	{object}	string
//	@Router			/users/{id} [put]
func (c Controller) UpdateUser(ctx *gin.Context) {
	var u UserUpdate
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user := c.service.UpdateUser(ctx, ctx.Param("id"), u)
	ctx.JSON(200, user)
}
