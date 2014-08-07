package resources

import (
	"encoding/json"
	"net/http"
)

// ResourceInterface defines an interface implementers of which
// should handle GET, POST requests from user.
type ResourceInterface interface {
	Dehydrate(w http.ResponseWriter, r *http.Request)
	Hydrate(w http.ResponseWriter, r *http.Request)
	GetData() interface{}
	GetBaseUrl() string
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

// Dehydrate handles GET requests.
// Method's name is not obvious but I like it :)
func (res *GenericResource) Dehydrate(w http.ResponseWriter, r *http.Request) {
	data := res.GetData()
	json_string, err := json.Marshal(data)
	if err != nil {
		json_string = []byte("")
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
		json_string = []byte("")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

// GetBaseUrl returns resource's BaseUrl attribute.
func (res *GenericResource) GetBaseUrl() string {
	return res.BaseUrl
}

// GetBaseUrl returns resource's Data attribute.
func (res *GenericResource) GetData() interface{} {
	return res.Data
}
