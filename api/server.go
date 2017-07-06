package api

import (
	"github.com/gorilla/mux"
)

func Server() *mux.Router {
	s := mux.NewRouter()
	s.HandleFunc("/", index)
	s.HandleFunc("/challenge", challengeCreate).Methods("POST")
	s.HandleFunc("/challenge/{challenge}/try", challengeTry).Methods("GET")
	return s
}
