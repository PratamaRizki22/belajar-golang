// interfaces/controllers/task_controller.go
package controllers

import (
	"clean_architecture/usecases"
	"encoding/json"
	"net/http"
)

type TaskController struct {
	CreateTaskUseCase usecases.CreateTaskUseCase
	GetTasksUseCase   usecases.GetTasksUseCase
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskData struct {
		Title string `json:"title"`
	}
	json.NewDecoder(r.Body).Decode(&taskData)

	task, err := c.CreateTaskUseCase.Execute(taskData.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.GetTasksUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
