package interfaces

import (
	"github.com/keijumt/training-go/domain"
)

type TaskRepository struct{}

// TODO Store memory or sql

func (repo *TaskRepository) FindBy(userID, taskID int) (task domain.Task, err error) {
	return domain.Task{ID: taskID, Title: "hoge", IsCompleted: false}, nil
}
func (repo *TaskRepository) FindAll(userID int) (tasks domain.Tasks, err error) {
	return []domain.Task{domain.Task{ID: 1, Title: "hoge1", IsCompleted: false}, domain.Task{ID: 2, Title: "hoge2", IsCompleted: false}}, nil
}
func (repo *TaskRepository) Store(userID int, title string) (task domain.Task, err error) {
	return domain.Task{ID: 1, Title: title, IsCompleted: false}, nil
}
