package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type AppRegistry interface {
	Name() string
	Services(container *dig.Container)
	ApiRoutes(c *dig.Container, router *gin.Engine)
}
