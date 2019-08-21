package infra

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/keijumt/training-go/domain"
	"net/http"
	"strconv"

	"github.com/keijumt/training-go/usecase"
)

type Controller struct {
	LoadTaskUseCase  usecase.LoadTaskUseCase
	LoadTasksUseCase usecase.LoadTasksUseCase
	StoreTaskUseCase usecase.StoreTaskUseCase
}

func (controller *Controller) LoadTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])
	taskID, _ := strconv.Atoi(vars["taskID"])
	task, err := controller.LoadTaskUseCase.Handle(userID, taskID)

	if err != nil {
		// TODO ErrorHandling
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (controller *Controller) LoadTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])

	tasks, err := controller.LoadTasksUseCase.Handle(userID)

	if err != nil {
		// TODO ErrorHandling
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (controller *Controller) StoreTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		// TODO ErrorHandling
		return
	}

	json.NewEncoder(w).Encode(domain.Task{ID: 1, Title: r.Form.Get("title")})
}
