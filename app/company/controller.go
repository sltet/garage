package company

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service: service}
}

type ControllerInterface interface {
	FindAllCompanies(ctx *gin.Context)
}

// FindAllCompanies godoc
// @Summary find all companies
// @Schemes
// @Description find all companies
// @Tags company
// @Accept json
// @Produce json
// @Success 200 {array} Company
// @Router /companies [get]
func (c Controller) FindAllCompanies(ctx *gin.Context) {
	ctx.JSON(200, []Company{
		c.service.CreateCompany("Landry & Fils", "2285 rue desmarteau", "Apt 3"),
	})
}
