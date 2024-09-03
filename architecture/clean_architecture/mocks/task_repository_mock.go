// mocks/task_repository_mock.go
package mocks

import (
	"clean_architecture/entities"
)

// MockTaskRepository mensimulasikan perilaku dari TaskRepository.
type MockTaskRepository struct {
	Tasks []*entities.Task
	Err   error
}

func (m *MockTaskRepository) Save(task *entities.Task) error {
	return m.Err // Mengembalikan error jika ada
}

func (m *MockTaskRepository) GetAll() ([]*entities.Task, error) {
	return m.Tasks, m.Err // Mengembalikan daftar tugas dan error
}
