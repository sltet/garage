package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type ControllerInterface interface {
	FindAllAppointments(ctx *gin.Context)
}

// FindAllAppointments godoc
//
//	@Summary	find all appointments
//	@Schemes
//	@Description	find all appointments
//	@Tags			appointment
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	Appointment
//	@Router			/appointments [get]
func (c Controller) FindAllAppointments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []Appointment{{
		Entity:     core.NewEntity(),
		LocationID: "1",
		ServiceID:  "1",
	}})
}
