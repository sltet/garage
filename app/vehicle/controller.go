package vehicle

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type ControllerInterface interface {
	FindAllVehicles(ctx *gin.Context)
}

// FindAllVehicles godoc
//
//	@Summary	find all vehicles
//	@Schemes
//	@Description	find all vehicles
//	@Tags			vehicle
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	Vehicle
//	@Router			/vehicles [get]
func (c Controller) FindAllVehicles(ctx *gin.Context) {
	ctx.JSON(200, []Vehicle{
		NewVehicle("Toyota", "Yaris", 2012),
		NewVehicle("Toyota", "Matrix", 2015),
		NewVehicle("Ford", "Escape", 2012),
		NewVehicle("Toyota", "Rav4", 2022),
	})
}
