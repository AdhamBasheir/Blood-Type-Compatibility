package routes

import (
	"blood-type-compatibility/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.GET("/users", controllers.ListUsers)
}
