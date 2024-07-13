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
	FindById(ctx *gin.Context)
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

// FindById godoc
//
//	@Summary	find by id
//	@Schemes
//	@Description	find by id
//	@Param			id	path	string	true	"the user to create"
//
// @Tags			company
// @Accept			json
// @Produce		json
// @Success		200	{object}	Company
// @Router			/companies/{id} [get]
func (c Controller) FindById(ctx *gin.Context) {
	company, err := c.service.FindById(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, company)
}
