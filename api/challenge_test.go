package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChallengeCreate(t *testing.T) {
	body := strings.NewReader("{}")
	request, _ := http.NewRequest("POST", fmt.Sprintf("/challenge"), body)
	recorder := httptest.NewRecorder()
	Server().ServeHTTP(recorder, request)
	response := recorder.Result()
	respBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 201 {
		t.Fatalf("Challenge create error. Want: 201. Got: %d.\nResponse body: %s", response.StatusCode, string(respBody))
	}
}
