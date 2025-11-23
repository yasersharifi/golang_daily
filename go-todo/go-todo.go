package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks []Task

func main() {
	r := gin.Default()

	r.Use(TimeMiddleware(), MyLogger())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "pong"})
	})

	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.POST("/tasks", createTask)
		api.GET("/tasks", findAllTasks)
		api.GET("/tasks/:id", findOneTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}

	r.Run(":9003")
}

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

func createTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&tasks); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(tasks) + 1
	newTask.Title = "Test title" + strconv.Itoa(newTask.ID)
	newTask.Description = "Test description" + strconv.Itoa(newTask.ID)

	tasks = append(tasks, newTask)

	ctx.JSON(201, gin.H{"msg": "created", "data": newTask})
}

func findAllTasks(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": tasks})
}

func findOneTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	for _, t := range tasks {
		if t.ID == id {
			ctx.JSON(200, t)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}

func updateTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var updated Task
	if err := ctx.ShouldBindJSON(&tasks); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			updated.ID = t.ID

			tasks[i] = updated
			ctx.JSON(200, updated)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}

func deleteTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if ctx.ShouldBindJSON(&tasks); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var deleted Task
	for i, t := range tasks {
		if t.ID == id {
			deleted = t
			// delete from arr
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(200, gin.H{"msg": "deleted", "data": deleted})
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}
