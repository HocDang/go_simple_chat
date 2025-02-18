package controller

import (
	"net/http"

	"time"

	"w3s/go-backend/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var request domain.TaskRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Create a new task using the validated fields
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Domain:    request.Domain,
		HookURL:   request.HookURL,
		CreatedAt: time.Now(),
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (u *TaskController) Fetch(c *gin.Context) {

	tasks, err := u.TaskUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (u *TaskController) GetByID(c *gin.Context) {

	taskID := c.Param("id")

	task, err := u.TaskUsecase.GetByID(c, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
