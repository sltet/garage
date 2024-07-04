package core

type AppRegistry interface {
	Name() string
	//ServicesDefinition(container *di.Builder)
	ApiRoutes() []ApiRouteDefinition
}
