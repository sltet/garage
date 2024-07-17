package operation

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service}
}

type ControllerInterface interface {
	FindAllServiceOperations(ctx *gin.Context)
	FindServiceOperationById(ctx *gin.Context)
}

// FindAllServiceOperations godoc
//
//	@Summary	find all garage service operations
//	@Schemes
//	@Description	find all garage service operations
//	@Tags			operation
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	ServiceOperation
//	@Router			/service-operations [get]
func (c Controller) FindAllServiceOperations(ctx *gin.Context) {
	operations, err := c.service.FindAllOperations(ctx)
	if err != nil {
		ctx.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, operations)
}

// FindServiceOperationById godoc
//
//	@Summary	find all garage service operation by id
//	@Schemes
//	@Description	find all garage service operation by id
//	@Param			id	path	string	true	"service operation id"
//	@Tags			operation
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ServiceOperation
//	@Router			/service-operations/{id} [get]
func (c Controller) FindServiceOperationById(ctx *gin.Context) {
	operation, err := c.service.FindById(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(err.Code(), gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, operation)
}
