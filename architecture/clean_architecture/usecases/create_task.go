// usecases/create_task.go
package usecases

import "clean_architecture/entities"

type TaskRepository interface {
	Save(task *entities.Task) error
}

type CreateTaskUseCase struct {
	TaskRepo TaskRepository
}

func (uc *CreateTaskUseCase) Execute(title string) (*entities.Task, error) {
	task := &entities.Task{Title: title, Completed: false}
	err := uc.TaskRepo.Save(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
