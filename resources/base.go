package resources

import (
	"encoding/json"
	"net/http"
)

// ResourceInterface defines an interface implementers of which
// should handle GET, POST requests from user.
type ResourceInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleGET(w http.ResponseWriter, r *http.Request)
	HandlePOST(w http.ResponseWriter, r *http.Request)
	GetBaseUrl() (string, error)
}

// EngineInterface defines an interface that should be implemented
// for every database engine you want to use with EngineResource.
type EngineInterface interface {
	HandleGETData(r *http.Request) (interface{}, error)
	HandlePOSTData(r *http.Request) error
}

// EngineResource implements ResourceInterface and is intended to use
// with databases wrappers that are implemented using EngineInterface(Engine attr).
// HTTP requests are passed to the :Engine: to generate data or do changes
// in the database/something else.
//
// To create a resource for some DB engine Foo:
//   1. Implement EngineInterface - FooEngine.
//   2. Create an instance of EngineResource using your FooEngine and other
// 		required arguments; and return it.
// 	 Constructor function will simplify this process.
type EngineResource struct {
	BaseUrl string
	Engine EngineInterface
	ConnectionAddress string
	DatabaseName string
	TableName string
}

func (res *EngineResource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res.HandleGET(w, r)
	case "POST":
		res.HandlePOST(w, r)
	}
}

func (res *EngineResource) HandleGET(w http.ResponseWriter, r *http.Request) {
	data, err := res.Engine.HandleGETData(r)

	json_string, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal resource data.", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

func (res *EngineResource) HandlePOST(w http.ResponseWriter, r *http.Request) {
	json_data := make(map[string]interface{})

	// TODO: Should I parse some data before?

	if err := res.Engine.HandlePOSTData(r); err != nil {
		json_data["success"] = false
	} else {
		json_data["success"] = true
	}
	json_string, err := json.Marshal(json_data)
	if err != nil {
		http.Error(w, "Failed to marshal response.", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

func (res *EngineResource) GetBaseUrl() (string, error) {
	return res.BaseUrl, nil
}
