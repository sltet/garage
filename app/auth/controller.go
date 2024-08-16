package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	// TODO: randomize it
	oauthStateString = "steve"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service: service}
}

type ControllerInterface interface {
	HandleGoogleLogin(ctx *gin.Context)
	ValidateGoogleToken(ctx *gin.Context)
	HandleCallbackGoogleLogin(ctx *gin.Context)
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
	//user, err := c.service.AuthenticateGoogleToken(ctx, idToken)
	//if err != nil {
	//	ctx.JSON(err.Code(), gin.H{"error": err.Error()})
	//	return
	//}
	//ctx.JSON(http.StatusOK, user)
}

func (c Controller) HandleCallbackGoogleLogin(ctx *gin.Context) {
	c.service.HandleCallbackGoogleLogin(ctx)
}

func (c Controller) HandleGoogleLogin(ctx *gin.Context) {
	c.service.HandleGoogleLogin(ctx)
}
