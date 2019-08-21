package usecase

import (
	"github.com/keijumt/training-go/domain"
)

type LoadTaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *LoadTaskInteractor) Handle(userID, taskId int) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.FindBy(userID, taskId)
	return
}
