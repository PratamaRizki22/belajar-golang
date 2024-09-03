// usecases/get_tasks_test.go
package usecases_test

import (
	"clean_architecture/entities"
	"clean_architecture/mocks"
	"clean_architecture/usecases"
	"testing"
)

func TestGetTasksUseCase(t *testing.T) {
	// Siapkan mock repository dengan data tertentu
	mockRepo := &mocks.MockTaskRepository{
		Tasks: []*entities.Task{
			{ID: 1, Title: "Task 1", Completed: false},
			{ID: 2, Title: "Task 2", Completed: true},
		},
	}

	// Inisialisasi use case dengan mock repository
	uc := usecases.GetTasksUseCase{TaskRepo: mockRepo}

	// Eksekusi use case
	tasks, err := uc.Execute()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verifikasi hasil
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "Task 1" || tasks[1].Title != "Task 2" {
		t.Errorf("Unexpected task titles: %v", tasks)
	}
}
