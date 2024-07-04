package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/core"
)

type Registry struct{}

func (r Registry) Name() string {
	return "user"
}

func (r Registry) ApiRoutes() []core.ApiRouteDefinition {
	return []core.ApiRouteDefinition{
		{
			Method:  core.GET,
			Path:    "/users",
			Handler: findAllUsers,
		},
	}
}

// FindAllUsers godoc
// @Summary find all users
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200
// @Router /users [get]
func findAllUsers(ctx *gin.Context) {
	ctx.JSON(200, []User{
		NewUser("Steve", "Landry"),
	})
}
