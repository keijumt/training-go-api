// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/keijumt/training-go/infra"
	"github.com/keijumt/training-go/interfaces"
	"github.com/keijumt/training-go/usecase"
)

// Injectors from wire.go:

func InjectController() infra.Controller {
	taskRepository := interfaces.NewTaskRepository()
	loadTaskUseCase := usecase.NewLoadTaskUseCase(taskRepository)
	loadTasksUseCase := usecase.NewLoadTasksUseCase(taskRepository)
	storeTaskUseCase := usecase.NewStoreTaskUseCase(taskRepository)
	controller := infra.NewController(loadTaskUseCase, loadTasksUseCase, storeTaskUseCase)
	return controller
}

// wire.go:

var ProvideController = wire.NewSet(
	ProvideLoadTaskUseCase,
	ProvideLoadTasksUseCase,
	ProvideStoreTaskUseCase,
	ProvideTaskRepository, infra.NewController,
)

var ProvideLoadTaskUseCase = wire.NewSet(usecase.NewLoadTaskUseCase)

var ProvideLoadTasksUseCase = wire.NewSet(usecase.NewLoadTasksUseCase)

var ProvideStoreTaskUseCase = wire.NewSet(usecase.NewStoreTaskUseCase)

var ProvideTaskRepository = wire.NewSet(interfaces.NewTaskRepository)