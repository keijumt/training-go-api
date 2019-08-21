package infra

import (
	"github.com/gorilla/mux"
	"github.com/keijumt/training-go/interfaces"
	"github.com/keijumt/training-go/usecase"
)

func NewRouter(r *mux.Router) {
	// TODO Dependency Injection
	var taskRepository usecase.TaskRepository
	taskRepository = &interfaces.TaskRepository{}

	var loadTaskUseCase usecase.LoadTaskUseCase
	loadTaskUseCase = &usecase.LoadTaskInteractor{TaskRepository: taskRepository}

	var loadTasksUseCase usecase.LoadTasksUseCase
	loadTasksUseCase = &usecase.LoadTasksInteractor{TaskRepository: taskRepository}

	var storeTaskUseCase usecase.StoreTaskUseCase
	storeTaskUseCase = &usecase.StoreTaskInteractor{TaskRepository: taskRepository}

	controller := Controller{LoadTaskUseCase: loadTaskUseCase, LoadTasksUseCase: loadTasksUseCase, StoreTaskUseCase: storeTaskUseCase}

	r.HandleFunc("/users/{userID}/tasks/{taskID}", controller.LoadTask).Methods("GET")
	r.HandleFunc("/users/{userID}/tasks", controller.LoadTasks).Methods("GET")
	r.HandleFunc("/users/{userID}/tasks", controller.StoreTask).Methods("POST")
}
