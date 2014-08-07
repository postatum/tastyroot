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
func Register(resource resources.ResourceInterface) {
	url, err := resource.GetBaseUrl()
	if err != nil {
		panic("Failed to get resource url via GetBaseUrl")
	}
	http.Handle(url, resource)
}
