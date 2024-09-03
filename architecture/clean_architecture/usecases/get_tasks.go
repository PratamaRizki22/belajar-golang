// usecases/get_tasks.go
package usecases

import (
	"clean_architecture/entities"
	"clean_architecture/interfaces/repositories" // Impor antarmuka dari lokasi yang benar
)

type GetTasksUseCase struct {
	TaskRepo repositories.TaskRepository // Gunakan antarmuka dari package repositories
}

func (uc *GetTasksUseCase) Execute() ([]*entities.Task, error) {
	return uc.TaskRepo.GetAll() // Pastikan TaskRepo mendefinisikan metode GetAll
}
