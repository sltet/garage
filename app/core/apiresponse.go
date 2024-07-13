package core

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	error string `json:"error,omitempty"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
