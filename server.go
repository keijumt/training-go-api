package main

import (
	"github.com/gorilla/mux"
	"github.com/keijumt/training-go/infra"
	"net/http"
)

func main() {
	// TODO Dependency Injection
	r := mux.NewRouter()
	infra.NewRouter(r)
	http.ListenAndServe("localhost:8080", r)
}
