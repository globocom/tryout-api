package api

import (
	"github.com/gorilla/mux"
)

func Server() *mux.Router {
	s := mux.NewRouter()
	s.HandleFunc("/", index)
	s.HandleFunc("/challenge", challengeCreate).Methods("POST")
	s.HandleFunc("/challenge/{challenge}/try", challengeTry).Methods("GET")
	s.HandleFunc("/challenge/{challenge}/{repo}/step", repoStepRegister).Methods("POST")
	s.HandleFunc("/challenge/{challenge}/{repo}", repoTryouts).Methods("GET")
	return s
}
