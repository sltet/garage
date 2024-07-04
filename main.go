package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sltet/garage/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		pong(ctx)
	})

	handler := ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	router.GET("/swagger/*any", handler)

	router.Run() // Listen and serve on 0.0.0.0:8080
}
