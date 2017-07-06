package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"../challenge"
	"../db"
)

func TestChallengeCreate(t *testing.T) {
	db.DialDB()
	defer db.DbSession.Close()
	coll := db.Coll("challenge")
	coll.Remove(bson.M{"_id": "xxx"})
	body := strings.NewReader(`{"name": "xxx", "paths": []}`)
	request, _ := http.NewRequest("POST", fmt.Sprintf("/challenge"), body)
	recorder := httptest.NewRecorder()
	Server().ServeHTTP(recorder, request)
	response := recorder.Result()
	respBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 201 {
		t.Fatalf("Challenge create error. Want: 201. Got: %d.\nResponse body: %s", response.StatusCode, string(respBody))
	}
	var c challenge.Challenge
	coll.Find(bson.M{"_id": "xxx"}).One(&c)
	if c.Name == "" {
		t.Fatal("challenge not saved. Want: xxx. Got: nil")
	}
	if string(respBody) != c.URL() {
		t.Fatalf("Wrong URL for challenge. Want: %s. Got: %s", c.URL(), respBody)
	}
}
