package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Registry struct{}

func (r Registry) Name() string {
	return "user"
}

func (r Registry) ServicesDefinition(c *dig.Container) {
	core.PanicOnError(c.Provide(NewController, dig.As(new(ControllerInterface))))
	core.PanicOnError(c.Provide(NewFactory, dig.As(new(FactoryInterface))))
	core.PanicOnError(c.Provide(NewService, dig.As(new(ServiceInterface))))
	core.PanicOnError(c.Provide(NewRepository, dig.As(new(RepositoryInterface))))
}

func (r Registry) SqlSchemaMigration(db *gorm.DB) {
	db.AutoMigrate(&User{})
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
			Path:   "/users",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).FindAllUsers(ctx)
			},
		},
		{
			Method: core.GET,
			Path:   "/users/:id",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).FindById(ctx)
			},
		},
		{
			Method: core.POST,
			Path:   "/users",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).CreateUser(ctx)
			},
		},
		{
			Method: core.PUT,
			Path:   "/users/:id",
			Handler: func(ctx *gin.Context, c *dig.Container) {
				controller(c).UpdateUser(ctx)
			},
		},
	}
}
