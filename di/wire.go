//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/keijumt/training-go/infra"
	"github.com/keijumt/training-go/interfaces"
	"github.com/keijumt/training-go/usecase"
)

var ProvideController = wire.NewSet(
	ProvideLoadTaskUseCase,
	ProvideLoadTasksUseCase,
	ProvideStoreTaskUseCase,
	ProvideTaskRepository,
	infra.NewController,
)

var ProvideLoadTaskUseCase = wire.NewSet(
	usecase.NewLoadTaskUseCase,
)

var ProvideLoadTasksUseCase = wire.NewSet(
	usecase.NewLoadTasksUseCase,
)

var ProvideStoreTaskUseCase = wire.NewSet(
	usecase.NewStoreTaskUseCase,
)

var ProvideTaskRepository = wire.NewSet(
	interfaces.NewTaskRepository,
)

func InjectController() infra.Controller {
	wire.Build(
		ProvideController,
	)
	return infra.Controller{}
}
