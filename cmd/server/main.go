package main

import (
	"TaskManagementSystem/internal/repositories"
	"TaskManagementSystem/internal/domains/dao"
	"TaskManagementSystem/internal/handlers"
	"TaskManagementSystem/internal/routers"
	"TaskManagementSystem/internal/services"

)

func main(){
	postgresvalues := &repositories.PostgresValue{ServerName:"localhost",Port:"5432",User:"postgres",Password:"postgres",DBName:"taskdb"}
	db := postgresvalues.ConnectToProstgresDB()
	taskDAO := &dao.TaskDAO{DB: db}
	taskServiceImpl  := &services.TaskServiceImpl{TaskDAO:taskDAO}
	taskHandler := &handlers.TaskHandler{Service:taskServiceImpl}
	taskRouter := routers.SetupTaskRouter(taskHandler)
	taskRouter.Run(":8080")
}