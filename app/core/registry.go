package core

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
)

type AppRegistry interface {
	Name() string
	ServicesDefinition(container *dig.Container)
	ApiRouteDefinitions() []ApiRouteDefinition
	RegisterCustomValidations(validator *validator.Validate)
}
