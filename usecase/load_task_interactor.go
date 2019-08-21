package usecase

import (
	"github.com/keijumt/training-go/domain"
)

type LoadTaskInteractor struct {
	TaskRepository TaskRepository
}

func NewLoadTaskUseCase(repository TaskRepository) LoadTaskUseCase {
	return &LoadTaskInteractor{repository}
}

func (interactor *LoadTaskInteractor) Handle(userID, taskId int) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.FindBy(userID, taskId)
	return
}
