package operation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service ServiceInterface
}

func NewController(service ServiceInterface) *Controller {
	return &Controller{service}
}

type ControllerInterface interface {
	FindAllOperations(ctx *gin.Context)
}

// FindAllOperations godoc
//
//	@Summary	find all garage operations
//	@Schemes
//	@Description	find all garage operations
//	@Tags			operation
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	ServiceOperation
//	@Router			/operations [get]
func (c Controller) FindAllOperations(ctx *gin.Context) {
	operations, err := c.service.FindAllOperations(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, operations)
}
