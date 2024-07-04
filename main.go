package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/company"
	"github.com/sltet/garage/core"
	_ "github.com/sltet/garage/docs"
	"github.com/sltet/garage/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /

func getRegistries() []core.AppRegistry {
	return []core.AppRegistry{
		user.Registry{},
		company.Registry{},
	}
}

func registerApiRoutes(router *gin.Engine) {
	for _, registry := range getRegistries() {
		for _, apiRoute := range registry.ApiRoutes() {
			if apiRoute.Method == core.GET {
				router.GET(apiRoute.Path, apiRoute.Handler)
				continue
			}
			if apiRoute.Method == core.POST {
				router.POST(apiRoute.Path, apiRoute.Handler)
				continue
			}
			if apiRoute.Method == core.DELETE {
				router.DELETE(apiRoute.Path, apiRoute.Handler)
				continue
			}
		}
	}
}

func main() {
	router := gin.Default()

	registerApiRoutes(router)

	handler := ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))
	router.GET("/swagger/*any", handler)

	router.Run() // Listen and serve on 0.0.0.0:8080
}
