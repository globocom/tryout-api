package api

import (
	"encoding/json"
	"fmt"
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
	err := db.Coll("repo").Find(bson.M{"name": vars["repo"]}).Sort("-_id").One(&repo)
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
	repo.Steps = steps
	repo.Version += 1
	err = db.Coll("repo").Insert(&repo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://tsuru.globoi.com/apps/%s", vars["challenge"]), nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer acbea31b589a270ec856569d9f0b6c88c23bb6a96c66ac5bcb1f7f54d12b1d69"))
	c := http.Client{}
	_, err = c.Do(req)
	w.WriteHeader(http.StatusOK)
}

func repoTryouts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%#v\n", vars)
	var repos []repository.Repository
	err := db.Coll("repo").Find(bson.M{"name": vars["repo"]}).Sort("-_id").All(&repos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	reposJson, err := json.Marshal(repos)
	w.Write(reposJson)
}
