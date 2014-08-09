package resources

import (
	"net/http"
)

// ResourceInterface defines an interface implementers of which
// should handle GET, POST requests from user.
type ResourceInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleGET(w http.ResponseWriter, r *http.Request)
	HandlePOST(w http.ResponseWriter, r *http.Request)
	GetData() (interface{}, error)
	GetBaseUrl() (string, error)
}
