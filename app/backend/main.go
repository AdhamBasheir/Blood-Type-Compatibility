package main

import (
	"blood-type-compatibility/initializers"
	"blood-type-compatibility/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	routes.RegisterHealthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterBloodRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
