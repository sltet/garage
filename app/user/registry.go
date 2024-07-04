package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
)

type Registry struct{}

func (r Registry) Name() string {
	return "user"
}

func (r Registry) ServicesDefinition(c *dig.Container) {
	_ = c.Provide(NewController, dig.As(new(ControllerInterface)))
}

func (r Registry) ApiRouteDefinitions() []core.ApiRouteDefinition {
	return []core.ApiRouteDefinition{
		{
			Method: core.GET,
			Path:   "/users",
			Handler: func(ctx *gin.Context, handler interface{}) {
				handler.(ControllerInterface).FindAllUsers(ctx)
			},
		},
	}
}

func (r Registry) ApiRoutesRegistration(c *dig.Container, router *gin.Engine) {
	for _, apiRoute := range r.ApiRouteDefinitions() {
		router.Handle(apiRoute.Method.String(), apiRoute.Path, func(ctx *gin.Context) {
			_ = c.Invoke(func(handler ControllerInterface) {
				apiRoute.Handler(ctx, handler)
			})
		})
	}
}
