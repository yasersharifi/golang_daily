package routes

import (
	"go-todo/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(rg *gin.RouterGroup) {
	task := rg.Group("/tasks")

	{
		task.POST("/", handlers.CreateTask)
		task.GET("/", handlers.FindAllTasks)
		task.GET("/:id", handlers.FindAllTasks)
		task.PUT("/:id", handlers.FindOneTask)
		task.DELETE("/:id", handlers.RemoveTask)
	}
}
