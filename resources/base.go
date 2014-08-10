package resources

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ResourceInterface defines an interface implementers of which
// should handle GET, POST requests from user.
type ResourceInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleGET(w http.ResponseWriter, r *http.Request)
	HandlePOST(w http.ResponseWriter, r *http.Request)
	GetBaseUrl() (string, error)
	GetRequestedId(r *http.Request) (string, error)
}

// EngineInterface defines an interface that should be implemented
// for every database engine you want to use with EngineResource.
type EngineInterface interface {
	HandleGETData(requestedId string) (interface{}, error)
	HandlePOSTData(requestData map[string]interface{}) error
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
	Engine  EngineInterface
}

func (res *EngineResource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res.HandleGET(w, r)
	case "POST":
		res.HandlePOST(w, r)
	}
}

// HandleGET parses request path, passes it to Engine.HandleGETData, checks
// for error and finally returns Engine's response in JSON format.
func (res *EngineResource) HandleGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestedId, _ := res.GetRequestedId(r)
	responseData, err := res.Engine.HandleGETData(requestedId)
	if err != nil {
		responseData = map[string]string{"error": err.Error()}
	}

	responseJson, err := json.Marshal(responseData)
	if err != nil {
		responseData = map[string]string{"error": err.Error()}
		responseJson, _ = json.Marshal(responseData)
	}

	w.Write(responseJson)
}

// HandlePOST parses JSON data from request, passed it to Engine.HandlePOSTData to
// create a new object in the Engine and returns JSON response identifying result.
func (res *EngineResource) HandlePOST(w http.ResponseWriter, r *http.Request) {
	requestData := make(map[string]interface{})
	responseData := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	errorJSONResp := func(w http.ResponseWriter, responseData map[string]interface{}) {
		responseData["created"] = false
		responseJson, _ := json.Marshal(responseData)
		w.Write(responseJson)
	}

	rawRequestBody := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(rawRequestBody); err != nil {
		responseData["error"] = "Failed to read request body."
		errorJSONResp(w, responseData)
		return
	}
	if err := json.Unmarshal(rawRequestBody, &requestData); err != nil {
		responseData["error"] = "Failed to parse JSON request body."
		errorJSONResp(w, responseData)
		return
	}
	if err := res.Engine.HandlePOSTData(requestData); err != nil {
		responseData["error"] = err.Error()
		errorJSONResp(w, responseData)
		return
	}
	responseData["created"] = true
	responseJson, _ := json.Marshal(responseData)
	w.Write(responseJson)
}

func (res *EngineResource) GetBaseUrl() (string, error) {
	return res.BaseUrl, nil
}

// GetRequestedId parses request path and extracts object id that is being
// requested. Schema such `objects/ID` is assumed, where ID is the of an
// `object` that is being requested.
func (res *EngineResource) GetRequestedId(r *http.Request) (string, error) {
	baseUrl, _ := res.GetBaseUrl()
	clearUrl := strings.Replace(r.URL.Path, baseUrl, "", 1)
	if len(clearUrl) == 0 {
		return "", nil
	}
	splitted := strings.Split(clearUrl, "/")
	return splitted[0], nil
}
