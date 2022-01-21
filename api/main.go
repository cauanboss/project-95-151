package main

import (
	"api-test/interface/routing"
	"log"
	"net/http"
)

func handleRequests() {
	http.Handle("/", routing.GetRoutes())
	log.Fatal(http.ListenAndServe(":8101", nil))
}

func main() {
	handleRequests()
}
