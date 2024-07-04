package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AppRegistry interface {
	Name() string
	ServicesDefinition(container *dig.Container)
	ApiRouteDefinitions() []ApiRouteDefinition
	ApiRoutesRegistration(c *dig.Container, router *gin.Engine)
}
