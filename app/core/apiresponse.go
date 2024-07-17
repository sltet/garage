package core

import "github.com/gin-gonic/gin"

type ApiError struct {
	code  int
	error string `json:"error,omitempty"`
}

func NewApiError(ctx *gin.Context, err DetailedError) {
	ctx.JSON(err.Code(), gin.H{"error": err.Error()})
}
