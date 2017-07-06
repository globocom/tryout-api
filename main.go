package main

import (
	"net/http"

	"./api"
)

func main() {
	http.ListenAndServe(":8080", api.Server())
}
