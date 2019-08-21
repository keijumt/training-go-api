package infra

import (
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, controller Controller) {
	r.HandleFunc("/users/{userID}/tasks/{taskID}", controller.LoadTask).Methods("GET")
	r.HandleFunc("/users/{userID}/tasks", controller.LoadTasks).Methods("GET")
	r.HandleFunc("/users/{userID}/tasks", controller.StoreTask).Methods("POST")
}
