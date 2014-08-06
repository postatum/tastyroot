package resources

import (
	"encoding/json"
	"net/http"
)

type JsonResponse map[string]interface{}

type ResourceInterface interface {
	Dehydrate(w http.ResponseWriter, r *http.Request)
	Hydrate(w http.ResponseWriter, r *http.Request)
	PrepareObject() map[string]interface{}
	GetUrl() string
}

type SimpleObjResource struct {
	Object interface{}
	Url    string
}

func (res *SimpleObjResource) Dehydrate(w http.ResponseWriter, r *http.Request) {
	var json_string []byte
	data := res.PrepareObject()
	json_string, err := json.Marshal(data)
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

func (res *SimpleObjResource) PrepareObject() map[string]interface{} {
	data := make(map[string]interface{})
	// TODO. This should iterate over a list of predifined keys
	// 		 not the interface{}
	for key, val := range res.Object {
		data[string(key)] = val
	}
	return data
}
