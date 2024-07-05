package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
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
// @Summary find all appointments
// @Schemes
// @Description find all appointments
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} Appointment
// @Router /appointments [get]
func (c Controller) FindAllAppointments(ctx *gin.Context) {
	ctx.JSON(200, []Appointment{{
		ID:         core.GetTimeBasedUUID().String(),
		LocationID: "1",
		ServiceID:  "1",
	}})
}
