package main

import (
	"net/http"

	"./api"
	"./db"
)

func main() {
	db.DialDB()
	defer db.DbSession.Close()
	http.ListenAndServe(":8080", api.Server())
}
