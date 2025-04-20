package routers

import (
	"TaskManagementSystem/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupTaskRouter(taskHandler *handlers.TaskHandler) *gin.Engine {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetAllTasks)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.POST("", taskHandler.CreateTask)
		tasks.PUT("", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}
	return router
}