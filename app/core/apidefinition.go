package core

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

const (
	GET    ApiMethod = "GET"
	POST   ApiMethod = "POST"
	PUT    ApiMethod = "PUT"
	PATCH  ApiMethod = "PATCH"
	DELETE ApiMethod = "DELETE"
)

type ApiMethod string
type ApiHandler func(ctx *gin.Context, c *dig.Container)

func (i ApiMethod) String() string {
	return string(i)
}

type ApiRouteDefinition struct {
	Method  ApiMethod
	Path    string
	Handler func(ctx *gin.Context, c *dig.Container)
}
