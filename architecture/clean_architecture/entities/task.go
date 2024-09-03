// entities/task.go
package entities

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Task) ToggleComplete() {
	t.Completed = !t.Completed
}
