// main.go
package main

import (
	"clean_architecture/constants"
	"clean_architecture/infrastructure/db"
	"clean_architecture/infrastructure/web"
	"clean_architecture/interfaces/controllers"
	"clean_architecture/usecases"
	"net/http"
)

func main() {
	taskRepo := db.NewInMemoryTaskRepository()
	createTaskUseCase := usecases.CreateTaskUseCase{TaskRepo: taskRepo}
	getTasksUseCase := usecases.GetTasksUseCase{TaskRepo: taskRepo}
	taskController := controllers.TaskController{
		CreateTaskUseCase: createTaskUseCase,
		GetTasksUseCase:   getTasksUseCase,
	}
	router := web.NewRouter(taskController)

	http.ListenAndServe(constants.ServerPort, router)
}
