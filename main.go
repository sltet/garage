package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sltet/garage/app/appointment"
	"github.com/sltet/garage/app/company"
	"github.com/sltet/garage/app/core"
	"github.com/sltet/garage/app/servicerequest"
	"github.com/sltet/garage/app/user"
	"github.com/sltet/garage/app/vehicle"
	_ "github.com/sltet/garage/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

// @BasePath /

func getRegistries() []core.AppRegistry {
	return []core.AppRegistry{
		user.Registry{},
		company.Registry{},
		vehicle.Registry{},
		appointment.Registry{},
		servicerequest.Registry{},
	}
}

func registerApiRoutes(c *dig.Container, router *gin.Engine) {
	for _, registry := range getRegistries() {
		for _, apiRoute := range registry.ApiRouteDefinitions() {
			router.Handle(apiRoute.Method.String(), apiRoute.Path, NewApiHandler(apiRoute.Handler, c))
		}
	}
}

func registerServices(ctn *dig.Container) {
	for _, registry := range getRegistries() {
		registry.ServicesDefinition(ctn)
	}
}

func NewApiHandler(apiHandler core.ApiHandler, c *dig.Container) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiHandler(ctx, c)
	}
}

func registerValidations() {
	for _, registry := range getRegistries() {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			registry.RegisterCustomValidations(v)
		}
	}
}

func main() {
	router := gin.Default()
	ctn := dig.New()

	registerServices(ctn)
	registerApiRoutes(ctn, router)
	registerValidations()

	router.Use()
	handler := ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))
	router.GET("/swagger/*any", handler)

	router.Run() // Listen and serve on 0.0.0.0:8080
}
