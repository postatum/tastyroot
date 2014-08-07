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
	cats := []Cat{
		{"Batman", 13, true},
		{"Banana", 3, true},
		{"Pong", 22, false},
	}
	cat := Cat{"Batman", 13, true}

	cat_resource := resources.GenericResource{cat, "/cat"}
	cats_resource := resources.GenericResource{cats, "/cats"}

	api.Register(&cat_resource)
	api.Register(&cats_resource)

	http.ListenAndServe(":8000", nil)
}