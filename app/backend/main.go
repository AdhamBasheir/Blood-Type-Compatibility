package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/middlewares"
	"blood-type-compatibility/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitLogger()
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	router := gin.New()
	router.Use(
		middlewares.Recovery(),
		middlewares.Logger(),
	)

	routes.RegisterHealthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterBloodRoutes(router)

	router.Run()
}
