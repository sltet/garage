package company

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/core"
	"go.uber.org/dig"
)

type Registry struct{}

func (r Registry) Name() string {
	return "company"
}

func (r Registry) Services(c *dig.Container) {
	_ = c.Provide(NewController, dig.As(new(ControllerInterface)))
}

func (r Registry) GetApiRouteDefinitions() []core.ApiRouteDefinition {
	return []core.ApiRouteDefinition{
		{
			Method: core.GET,
			Path:   "/companies",
			Handler: func(ctx *gin.Context, handler interface{}) {
				handler.(ControllerInterface).FindAllCompanies(ctx)
			},
		},
	}
}

func (r Registry) ApiRoutes(c *dig.Container, router *gin.Engine) {
	for _, apiRoute := range r.GetApiRouteDefinitions() {
		router.Handle(apiRoute.Method.String(), apiRoute.Path, func(ctx *gin.Context) {
			_ = c.Invoke(func(handler ControllerInterface) {
				apiRoute.Handler(ctx, handler)
			})
		})
	}
}
