package company

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/core"
)

type Registry struct{}

func (r Registry) Name() string {
	return "company"
}

func (r Registry) ApiRoutes() []core.ApiRouteDefinition {
	return []core.ApiRouteDefinition{
		{
			Method:  core.GET,
			Path:    "/companies",
			Handler: findAllCompanies,
		},
	}
}

// findAllCompanies godoc
// @Summary find all users
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200
// @Router /companies [get]
func findAllCompanies(ctx *gin.Context) {
	ctx.JSON(200, []Company{{
		Name:         "Landry & Fils",
		AddressLine1: "2285 rue desmarteau",
		AddressLine2: "Apt 3",
	}})
}
