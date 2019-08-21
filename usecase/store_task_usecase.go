package usecase

import "github.com/keijumt/training-go/domain"

type StoreTaskUseCase interface {
	Handle(userID int, title string) (task domain.Task, err error)
}
