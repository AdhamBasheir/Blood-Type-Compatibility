package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/middlewares"
	"blood-type-compatibility/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
	initializers.InitLogger()
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
