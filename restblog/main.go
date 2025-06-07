package main

import (
	"log"
	"net/http"
	route "restblog/route"
)

func main() {
	r := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
