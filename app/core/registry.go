package core

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type AppRegistry interface {
	Name() string
	ServicesDefinition(container *dig.Container)
	SqlSchemaMigration(db *gorm.DB)
	ApiRouteDefinitions() []ApiRouteDefinition
	RegisterCustomValidations(validator *validator.Validate)
}
