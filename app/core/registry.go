package core

import (
	"go.uber.org/dig"
)

type AppRegistry interface {
	Name() string
	ServicesDefinition(container *dig.Container)
	ApiRouteDefinitions() []ApiRouteDefinition
}
