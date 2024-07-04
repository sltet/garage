package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/company"
	"github.com/sltet/garage/core"
	_ "github.com/sltet/garage/docs"
	"github.com/sltet/garage/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

// @BasePath /

func getRegistries() []core.AppRegistry {
	return []core.AppRegistry{
		user.Registry{},
		company.Registry{},
	}
}

func registerApiRoutes(ctn *dig.Container, router *gin.Engine) {
	for _, registry := range getRegistries() {
		registry.ApiRoutesRegistration(ctn, router)
	}
}

func registerServices(ctn *dig.Container) {
	for _, registry := range getRegistries() {
		registry.ServicesDefinition(ctn)
	}
}

func main() {
	router := gin.Default()
	ctn := dig.New()

	registerServices(ctn)
	registerApiRoutes(ctn, router)

	handler := ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))
	router.GET("/swagger/*any", handler)

	router.Run() // Listen and serve on 0.0.0.0:8080
}
