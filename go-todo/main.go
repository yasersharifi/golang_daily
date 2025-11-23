package main

import (
	"go-todo/internal/middlewares"
	"go-todo/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.TimeMiddleware(), middlewares.MyLogger())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "pong"})
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")

	routes.TaskRoutes(v1)

	r.Run(":9003")
}
