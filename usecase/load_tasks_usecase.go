package usecase

import "github.com/keijumt/training-go/domain"

type LoadTasksUseCase interface {
	Handle(userID int) (domain.Tasks, error)
}
