// usecases/get_tasks_test.go
package usecases_test

import (
	"clean_architecture/entities"
	"clean_architecture/infrastructure/db"
	"clean_architecture/usecases"
	"testing"
)

func TestGetTasksUseCase_Execute(t *testing.T) {
	repo := db.NewInMemoryTaskRepository()
	useCase := usecases.GetTasksUseCase{TaskRepo: repo}

	// Setup data
	task1 := &entities.Task{Title: "Task 1"}
	repo.Save(task1)

	// Execute use case
	tasks, err := useCase.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Validate results
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Task 1" {
		t.Fatalf("expected task title to be 'Task 1', got '%s'", tasks[0].Title)
	}
}
