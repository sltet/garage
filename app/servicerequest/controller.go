package servicerequest

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type ControllerInterface interface {
	FindAllServices(ctx *gin.Context)
}

// FindAllServices godoc
//
//	@Summary	find all service requests
//	@Schemes
//	@Description	find all services requests
//	@Tags			service request
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	ServiceRequest
//	@Router			/service-requests [get]
func (c Controller) FindAllServices(ctx *gin.Context) {
	ctx.JSON(200, []ServiceRequest{{
		Name: "Mechanic",
	}})
}
