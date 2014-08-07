package api

import (
	"net/http"
	"tastyroot/resources"
)

// Register func registers an implementation of ResourceInterface interface
// to handle requests at res.BaseUrl.
//
// Example:
//
// import "tastyroot/api"
//
// api.Register(&someResource)
//
func Register(res resources.ResourceInterface) {
	http.HandleFunc(res.GetBaseUrl(), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			res.Dehydrate(w, r)
		case "POST":
			res.Hydrate(w, r)
		}
	})
}
