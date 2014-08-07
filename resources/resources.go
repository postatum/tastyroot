package resources

import (
	"encoding/json"
	"net/http"
)

// ResourceInterface defines an interface implementers of which
// should handle GET, POST requests from user.
type ResourceInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Dehydrate(w http.ResponseWriter, r *http.Request)
	Hydrate(w http.ResponseWriter, r *http.Request)
	GetData() (interface{}, error)
	GetBaseUrl() (string, error)
}

// GenericResource implements ResourceInterface and allows to serve
// JSON data of a given :Data:. Marshal fields visibility rules apply.
//
// Example:
//
// import (
// 	"tastyroot/resources"
// 	"tastyroot/api"
// )
//
// type Cat struct {
// 	Name  string
// 	Age   int
// }
//
// bane = Cat{"Bane", 13}
//
// cat_resource := resources.GenericResource{cat, "/cat"}
//
// api.Register(&cat_resource)

type GenericResource struct {
	Data    interface{}
	BaseUrl string
}

// ServeHTTP checks the method and calls certain method to handle the request.
func (res *GenericResource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res.Dehydrate(w, r)
	case "POST":
		res.Hydrate(w, r)
	}
}

// Dehydrate handles GET requests.
// Method's name is not obvious but I like it :)
func (res *GenericResource) Dehydrate(w http.ResponseWriter, r *http.Request) {
	data, err := res.GetData()
	if err != nil {
		http.Error(w, "Failed to get resource data.", 500)
		return
	}
	json_string, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal resource data.", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

// Hydrate handles only POST (for now) requests and does nothing useful (for now).
func (res *GenericResource) Hydrate(w http.ResponseWriter, r *http.Request) {
	test_response := make(map[string]interface{})
	test_response["ping"] = "pong"
	json_string, err := json.Marshal(test_response)
	if err != nil {
		http.Error(w, "Failed to process the input.", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

// GetBaseUrl returns resource's BaseUrl attribute.
func (res *GenericResource) GetBaseUrl() (string, error) {
	return res.BaseUrl, nil
}

// GetBaseUrl returns resource's Data attribute.
func (res *GenericResource) GetData() (interface{}, error) {
	return res.Data, nil
}
