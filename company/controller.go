package company

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
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
// @Success 200 {object} Company
// @Router /companies [get]
func (c Controller) FindAllCompanies(ctx *gin.Context) {
	ctx.JSON(200, []Company{{
		Name:         "Landry & Fils",
		AddressLine1: "2285 rue desmarteau",
		AddressLine2: "Apt 3",
	}})
}
