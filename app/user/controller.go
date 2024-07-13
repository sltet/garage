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
	FindById(ctx *gin.Context)
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

// FindById godoc
//
//	@Summary	find user by id
//	@Schemes
//	@Description	find user by id
//	@Param			id	path	string	true	"user id"
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Router			/users/{id} [get]
func (c Controller) FindById(ctx *gin.Context) {
	user, err := c.service.FindById(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
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
		return
	}
	user, err := c.service.CreateUser(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
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
		return
	}
	user, err := c.service.UpdateUser(ctx, ctx.Param("id"), u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
