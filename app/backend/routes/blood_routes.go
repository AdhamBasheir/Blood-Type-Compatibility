package routes

import (
	"blood-type-compatibility/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBloodRoutes(r *gin.Engine) {
	r.GET("/compatible-donors", controllers.GetCompatibleDonors)
}
