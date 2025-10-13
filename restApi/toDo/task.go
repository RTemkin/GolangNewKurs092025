package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title, descrittion string) Task {
	return Task{
		Title:       title,
		Description: descrittion,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	doneTime := time.Now()

	t.Completed = true
	t.CompletedAt = &doneTime
}
