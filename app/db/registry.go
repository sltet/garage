package db

import (
	"github.com/go-playground/validator/v10"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Registry struct{}

func (r Registry) Name() string {
	return "db"
}

func (r Registry) ServicesDefinition(c *dig.Container) {
	core.PanicOnError(c.Provide(NewDatabase, dig.As(new(DatabaseInterface))))
}

func (r Registry) SqlSchemaMigration(db *gorm.DB) {}

func (r Registry) RegisterCustomValidations(validator *validator.Validate) {}

func (r Registry) ApiRouteDefinitions() []core.ApiRouteDefinition {
	return []core.ApiRouteDefinition{}
}
