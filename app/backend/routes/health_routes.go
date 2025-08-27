package routes

import (
	"blood-type-compatibility/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/healthz", controllers.Health)
}
