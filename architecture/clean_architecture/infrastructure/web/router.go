// infrastructure/web/router.go
package web

import (
	"clean_architecture/interfaces/controllers"
	"net/http"
)

func NewRouter(taskController controllers.TaskController) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/tasks", taskController.CreateTask)
	router.HandleFunc("/tasks/list", taskController.GetTasks)
	return router
}
