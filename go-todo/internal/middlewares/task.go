package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start)
		fmt.Println("Duration: ", duration)
	}
}

func MyLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Path: ", ctx.Request.URL.Path)
		ctx.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Token")

		// Also check the token not expired
		if token == "" {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
