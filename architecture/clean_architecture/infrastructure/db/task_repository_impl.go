// infrastructure/db/task_repository_impl.go
package db

import (
	"clean_architecture/entities"
	"sync"
)

type InMemoryTaskRepository struct {
	tasks []*entities.Task
	mu    sync.Mutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: []*entities.Task{},
	}
}

func (r *InMemoryTaskRepository) Save(task *entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, task)
	return nil
}

// Implementasi metode GetAll
func (r *InMemoryTaskRepository) GetAll() ([]*entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.tasks, nil
}
