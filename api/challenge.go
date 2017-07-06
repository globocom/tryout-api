package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"

	"../challenge"
	"../db"
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
	w.Write([]byte(c.URL()))
}

func challengeTry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var c challenge.Challenge
	err := db.Coll("challenge").Find(bson.M{"_id": vars["challenge"]}).One(&c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("starting tryout.\n"))
	err = c.Start(r.URL.Query().Get("repo"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
