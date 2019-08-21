package usecase

import (
	"github.com/keijumt/training-go/domain"
)

type LoadTaskUseCase interface {
	Handle(userID, taskID int) (domain.Task, error)
}
