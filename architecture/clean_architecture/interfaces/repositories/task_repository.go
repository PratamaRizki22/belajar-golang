// interfaces/repositories/task_repository.go
package repositories

import "clean_architecture/entities"

// Antarmuka yang benar untuk TaskRepository
type TaskRepository interface {
	Save(task *entities.Task) error
	GetAll() ([]*entities.Task, error) // Metode GetAll harus didefinisikan di sini
}
