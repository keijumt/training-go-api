package usecase

import (
	"github.com/keijumt/training-go/domain"
)

type TaskRepository interface {
	FindBy(userID, taskID int) (domain.Task, error)
	FindAll(userID int) (domain.Tasks, error)
	Store(userID int, title string) (domain.Task, error)
}
