package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../db"
	"../repository"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func repoStepRegister(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var repo repository.Repository
	err := db.Coll("repo").Find(bson.M{"_id": vars["repo"]}).One(&repo)
	if err != nil {
		if err == mgo.ErrNotFound {
			repo = repository.Repository{
				Name:      vars["repo"],
				Challenge: vars["challenge"],
				Version:   0,
				Steps:     []repository.Step{},
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
	var steps []repository.Step
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, &steps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	repo.Steps = append(repo.Steps, steps...)
	repo.Version += 1
	_, err = db.Coll("repo").Upsert(bson.M{"_id": repo.Name}, &repo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
