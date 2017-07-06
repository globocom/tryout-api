package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../challenge"
)

func challengeCreate(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var c challenge.Challenge
	err := dec.Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error parsing json: %s", err.Error())))
		return
	}
	err = challenge.Create(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error creating challenge: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("challenge created: %#v", c)))
}
