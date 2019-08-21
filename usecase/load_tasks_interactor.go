package usecase

import "github.com/keijumt/training-go/domain"

type LoadTasksInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *LoadTasksInteractor) Handle(userID int) (tasks domain.Tasks, err error) {
	tasks, err = interactor.TaskRepository.FindAll(userID)
	return
}
