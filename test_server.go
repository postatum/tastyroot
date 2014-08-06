package main

import (
	"net/http"
	"tastyroot/api"
	"tastyroot/resources"
)

type Cat struct {
	Name  string
	Age   int
	Alive bool
}

func main() {
	cat := Cat{"Batman", 13, true}
	resource := resources.SimpleObjResource{cat, "/cat"}
	api.Register(&resource)

	http.ListenAndServe(":8000", nil)
}
