package handlers

import (
	"TaskManagementSystem/internal/constants"
	"TaskManagementSystem/internal/models/postgres1"
	"TaskManagementSystem/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	Service services.TaskService
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	pageParam := c.DefaultQuery("page", "1")
	pageSizeParam := c.DefaultQuery("page_size", "10")
	status := c.Query("status")

	page, err1 := strconv.Atoi(pageParam)
	pageSize, err2 := strconv.Atoi(pageSizeParam)

	if err1 != nil || err2 != nil || page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    "Invalid pagination parameters",
		})
		return
	}

	tasks, err := h.Service.GetAllTasks(page, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.APIResponse{
			StatusCode: constants.UnknownError.GetResponseStatus(),
			Message:    constants.UnknownError.GetResponseMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, constants.APIResponse{
		StatusCode: constants.Success.GetResponseStatus(),
		Message:    constants.Success.GetResponseMessage(),
		Data:       tasks,
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    constants.InvalidRequest.GetResponseMessage(),
		})
		return
	}

	task, err := h.Service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, constants.APIResponse{
			StatusCode: constants.DataNotFound.GetResponseStatus(),
			Message:    constants.DataNotFound.GetResponseMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, constants.APIResponse{
		StatusCode: constants.Success.GetResponseStatus(),
		Message:    constants.Success.GetResponseMessage(),
		Data:       task,
	})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task postgres1.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    constants.InvalidRequest.GetResponseMessage(),
		})
		return
	}

	createdTask, err := h.Service.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.APIResponse{
			StatusCode: constants.UnknownError.GetResponseStatus(),
			Message:    constants.UnknownError.GetResponseMessage(),
		})
		return
	}

	c.JSON(http.StatusCreated, constants.APIResponse{
		StatusCode: constants.Success.GetResponseStatus(),
		Message:    constants.Success.GetResponseMessage(),
		Data:       createdTask,
	})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    constants.InvalidRequest.GetResponseMessage(),
		})
		return
	}

	var task postgres1.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    constants.InvalidRequest.GetResponseMessage(),
		})
		return
	}

	updatedTask, err := h.Service.UpdateTask(id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.APIResponse{
			StatusCode: constants.UnknownError.GetResponseStatus(),
			Message:    constants.UnknownError.GetResponseMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, constants.APIResponse{
		StatusCode: constants.Success.GetResponseStatus(),
		Message:    constants.Success.GetResponseMessage(),
		Data:       updatedTask,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, constants.APIResponse{
			StatusCode: constants.InvalidRequest.GetResponseStatus(),
			Message:    constants.InvalidRequest.GetResponseMessage(),
		})
		return
	}

	err = h.Service.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.APIResponse{
			StatusCode: constants.UnknownError.GetResponseStatus(),
			Message:    constants.UnknownError.GetResponseMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, constants.APIResponse{
		StatusCode: constants.Success.GetResponseStatus(),
		Message:    "Task deleted successfully",
	})
}
