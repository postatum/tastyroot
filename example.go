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

func SimpleResourceExample() {
	cats := []Cat{
		{"Batman", 13, true},
		{"Banana", 3, true},
		{"Pong", 22, false},
	}
	cat := Cat{"Batman", 13, true}

	cat_resource := &resources.SimpleResource{cat, "/cat"}
	cats_resource := &resources.SimpleResource{cats, "/cats"}

	api.Register(cat_resource)
	api.Register(cats_resource)
}

func MongoResourceExample() {
	cats_resource := resources.MongoResource(
		"/cats/",
		"127.0.0.1:27017",
		"godb",
		"cats",
		10,
	)
	api.Register(cats_resource)
}

func main() {
	// SimpleResourceExample()

	MongoResourceExample()

	http.ListenAndServe(":8000", nil)
}
