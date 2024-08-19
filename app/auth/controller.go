package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service: service}
}

type ControllerInterface interface {
	HandleLogin(ctx *gin.Context)
	ValidateGoogleToken(ctx *gin.Context)
	HandleCallback(ctx *gin.Context)
	HandleLogout(ctx *gin.Context)
}

// ValidateGoogleToken godoc
//
//	@Summary	Validate google token
//	@Schemes
//	@Description	Validate google token
//	@Param			token	query	string	true	"the google id token"
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Company
//	@Failure		400	{object}	string
//	@Router			/auth/google [get]
func (c Controller) ValidateGoogleToken(ctx *gin.Context) {
	idToken := ctx.Param("token")
	if idToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "token cannot be empty"})
		return
	}
}

func (c Controller) HandleCallback(ctx *gin.Context) {
	user, err := c.service.HandleCallback(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c Controller) HandleLogin(ctx *gin.Context) {
	c.service.HandleLogin(ctx)
}

func (c Controller) HandleLogout(ctx *gin.Context) {
	err := c.service.HandleLogout(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
