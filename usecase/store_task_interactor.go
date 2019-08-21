package usecase

import "github.com/keijumt/training-go/domain"

type StoreTaskInteractor struct {
	TaskRepository TaskRepository
}

func NewStoreTaskUseCase(repository TaskRepository) StoreTaskUseCase {
	return &StoreTaskInteractor{repository}
}

func (interactor *StoreTaskInteractor) Handle(userID int, title string) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.Store(userID, title)
	return
}
