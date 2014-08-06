package resources

import (
	"encoding/json"
	"net/http"
)

type JsonResponse map[string]interface{}

type ResourceInterface interface {
	Dehydrate(w http.ResponseWriter, r *http.Request)
	Hydrate(w http.ResponseWriter, r *http.Request)
	GetUrl() string
}

type SimpleObjResource struct {
	Object interface{}
	Url    string
}

func (res *SimpleObjResource) Dehydrate(w http.ResponseWriter, r *http.Request) {
	json_string, err := json.Marshal(res.Object)
	if err != nil {
		json_string = []byte("")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

func (res *SimpleObjResource) Hydrate(w http.ResponseWriter, r *http.Request) {
	test_response := make(map[string]interface{})
	test_response["ping"] = "pong"
	json_string, err := json.Marshal(test_response)
	if err != nil {
		json_string = []byte("")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_string)
}

func (res *SimpleObjResource) GetUrl() string {
	return res.Url
}
