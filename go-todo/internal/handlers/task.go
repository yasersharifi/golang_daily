package handlers

import (
	"go-todo/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Tasks = []models.Task{}

func CreateTask(ctx *gin.Context) {
	var input models.TaskInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newTask := models.Task{
		ID:          len(Tasks) + 1,
		Title:       input.Title,
		Description: input.Description,
		Done:        input.Done,
	}

	Tasks = append(Tasks, newTask)

	ctx.JSON(201, newTask)
}

func FindAllTasks(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": Tasks})
}

func FindOneTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	for _, t := range Tasks {
		if t.ID == id {
			ctx.JSON(200, t)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}

func UpdateTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var input models.TaskInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for i, t := range Tasks {
		if t.ID == id {
			updated := models.Task{
				ID:          t.ID,
				Title:       input.Title,
				Description: input.Description,
				Done:        input.Done,
			}
			Tasks[i] = updated
			ctx.JSON(200, updated)
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}

func RemoveTask(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			ctx.JSON(200, gin.H{"msg": "task deleted", "data": t})
			return
		}
	}

	ctx.JSON(404, gin.H{"error": "task not found"})
}
