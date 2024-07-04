package core

import "github.com/gin-gonic/gin"

const (
	GET    ApiMethod = "GET"
	POST   ApiMethod = "POST"
	DELETE ApiMethod = "DELETE"
)

type ApiMethod string

func (i ApiMethod) String() string {
	return string(i)
}

type ApiRouteDefinition struct {
	Method  ApiMethod
	Path    string
	Handler func(ctx *gin.Context)
}
