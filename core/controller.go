package core

import "github.com/gin-gonic/gin"

type CrudController interface {
	FindAll(ctx *gin.Context)
}
