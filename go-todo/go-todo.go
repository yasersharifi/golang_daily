package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	r := gin.Default()

	var tasks []Task

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "pong"})
	})

	r.POST("/tasks", func(ctx *gin.Context) {
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
	})

	r.GET("/tasks", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": tasks})
	})

	r.GET("/tasks/:id", func(ctx *gin.Context) {
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
	})

	r.PUT("/tasks/:id", func(ctx *gin.Context) {
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
	})

	r.DELETE("/tasks/:id", func(ctx *gin.Context) {
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
				ctx.JSON(200, gin.H{"msg": "deleted", "data": deleted})
				return
			}
		}

		ctx.JSON(404, gin.H{"error": "task not found"})
	})

	r.Run(":9003")
}
