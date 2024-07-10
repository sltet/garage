package company

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
	FindAllCompanies(ctx *gin.Context)
}

// FindAllCompanies godoc
//
//	@Summary	find all companies
//	@Schemes
//	@Description	find all companies
//	@Tags			company
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	Company
//	@Router			/companies [get]
func (c Controller) FindAllCompanies(ctx *gin.Context) {
	companies, err := c.service.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, companies)
}
