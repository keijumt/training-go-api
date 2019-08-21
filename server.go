package main

import (
	"github.com/gorilla/mux"
	"github.com/keijumt/training-go/di"
	"github.com/keijumt/training-go/infra"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	infra.NewRouter(r, di.InjectController())
	http.ListenAndServe("localhost:8080", r)
}
