package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Registry struct{}

func (r Registry) Name() string {
	return "auth"
}

func (r Registry) ServicesDefinition(c *dig.Container) {
	core.PanicOnError(c.Provide(NewController, dig.As(new(ControllerInterface))))
	core.PanicOnError(c.Provide(NewGoogleService, dig.As(new(GoogleServiceInterface))))
	core.PanicOnError(c.Provide(NewService, dig.As(new(ServiceInterface))))
}

func (r Registry) SqlSchemaMigration(db *gorm.DB) {
	//core.PanicOnError(db.AutoMigrate(&Company{}))
}

func (r Registry) RegisterCustomValidations(validator *validator.Validate) {

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
			Path:   "/auth/google/callback",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).HandleCallbackGoogleLogin(ctx)
			},
		},
		{
			Method: core.GET,
			Path:   "/auth/login/google",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).HandleGoogleLogin(ctx)
			},
		},
	}
}
