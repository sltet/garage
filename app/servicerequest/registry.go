package servicerequest

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
)

type Registry struct{}

func (r Registry) Name() string {
	return "company"
}

func (r Registry) ServicesDefinition(c *dig.Container) {
	core.PanicOnError(c.Provide(NewController, dig.As(new(ControllerInterface))))
}

func (r Registry) ApiRouteDefinitions() []core.ApiRouteDefinition {
	controller := func(c *dig.Container) (ctrl ControllerInterface) {
		core.PanicOnError(c.Invoke(func(handler ControllerInterface) {
			ctrl = handler
		}))
		return ctrl
	}

	return []core.ApiRouteDefinition{
		{
			Method: core.GET,
			Path:   "/services",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).FindAllServices(ctx)
			},
		},
	}
}
