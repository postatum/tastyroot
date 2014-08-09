package resources

import (
	"encoding/json"
	"net/http"
)

// SimpleResource implements ResourceInterface and allows to serve
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
// cat_resource := resources.SimpleResource{cat, "/cat"}
//
// api.Register(&cat_resource)

type SimpleResource struct {
	Data    interface{}
	BaseUrl string
}

// ServeHTTP checks the method and calls certain method to handle the request.
func (res *SimpleResource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res.HandleGET(w, r)
	case "POST":
		res.HandlePOST(w, r)
	}
}

// HandleGET handles GET requests.
// Method's name is not obvious but I like it :)
func (res *SimpleResource) HandleGET(w http.ResponseWriter, r *http.Request) {
	json_string, err := json.Marshal(res.Data)
	if err != nil {
		http.Error(w, "Failed to marshal resource data.", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

// HandlePOST handles only POST (for now) requests and does nothing useful (for now).
func (res *SimpleResource) HandlePOST(w http.ResponseWriter, r *http.Request) {
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
func (res *SimpleResource) GetBaseUrl() (string, error) {
	return res.BaseUrl, nil
}
