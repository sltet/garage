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
	//GetAuthLoginUrl(ctx *gin.Context)
	//Authorize(ctx *gin.Context)
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
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	user, err := c.service.HandleCallback(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c Controller) HandleLogin(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	c.service.HandleLogin(ctx)
}

// GetAuthLoginUrl godoc
//
//	@Summary	Get oauth login url
//	@Schemes
//	@Description	Get oauth login url
//	@Param			provider	path	string	true	"the login provider"
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/auth/:provider/login [get]
//func (c Controller) GetAuthLoginUrl(ctx *gin.Context) {
//	q := ctx.Request.URL.Query()
//	q.Add("provider", ctx.Param("provider"))
//	ctx.Request.URL.RawQuery = q.Encode()
//	url, err := c.service.GetAuthLoginUrl(ctx)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"oauth_url": url})
//	return
//}

// Authorize godoc
//
//	@Summary	Get oauth login url
//	@Schemes
//	@Description	Get oauth login url
//	@Param			code	query	string	true	"the oauth code"
//	@Param			state	query	string	true	"the state "
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/auth/:provider/authorize [get]
//func (c Controller) Authorize(ctx *gin.Context) {
//	q := ctx.Request.URL.Query()
//	q.Add("provider", ctx.Param("provider"))
//	ctx.Request.URL.RawQuery = q.Encode()
//	url, err := c.service.Authorize(ctx)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"oauth_url": url})
//	return
//}

func (c Controller) HandleLogout(ctx *gin.Context) {
	q := ctx.Request.URL.Query()
	q.Add("provider", ctx.Param("provider"))
	ctx.Request.URL.RawQuery = q.Encode()
	err := c.service.HandleLogout(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
