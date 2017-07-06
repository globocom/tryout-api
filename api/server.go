package api

import (
	"github.com/gorilla/mux"
)

func Server() *mux.Router {
	s := mux.NewRouter()
	s.HandleFunc("/", index)
	s.HandleFunc("/challenge", challengeCreate).Methods("POST")
	return s
}
