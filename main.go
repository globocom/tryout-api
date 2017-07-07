package main

import (
	"net/http"

	"./api"
	"./db"
)

func main() {
	err := db.DialDB()
	if err != nil {
		panic(err)
	}
	defer db.DbSession.Close()
	http.ListenAndServe("0.0.0.0:8888", api.Server())
}
