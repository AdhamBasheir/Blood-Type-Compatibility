package controllers

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func Health(ctx *gin.Context) {
	// Placeholder for a more comprehensive health check
	ctx.JSON(200, gin.H{
		"status": "healthy",
	})
}
