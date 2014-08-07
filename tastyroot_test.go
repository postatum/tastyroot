package main

import (
	"net/http"
	"net/http/httptest"
	"tastyroot/resources"
	"testing"
)

type TestCat struct {
	Name string
}

func TestGenericResourceDehydrate(t *testing.T) {
	testUrl := "/cat"
	resource := resources.GenericResource{
		TestCat{"Meow"},
		testUrl,
	}

	// Successful processing
	expectedBody := "{\"Name\":\"Meow\"}"
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	resource.Dehydrate(recorder, req)

	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("Body (%s) did not match expectation (%s).",
			recorder.Body.String(),
			expectedBody)
	}
}

func TestGenericResourceHydrate(t *testing.T) {
	testUrl := "/cat"
	resource := resources.GenericResource{
		TestCat{"Meow"},
		testUrl,
	}

	expectedBody := "{\"ping\":\"pong\"}"
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "http://example.com", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	resource.Hydrate(recorder, req)

	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("Body (%s) did not match expectation (%s).",
			recorder.Body.String(),
			expectedBody)
	}
}

func TestGenericResourceGetters(t *testing.T) {
	testUrl := "/cat2"
	cat := TestCat{"Meow"}
	resource := resources.GenericResource{cat, testUrl}

	if data, _ := resource.GetData(); data != cat {
		t.Errorf("GetData didn't return expected data")
	}
	if url, _ := resource.GetBaseUrl(); url != testUrl {
		t.Errorf("GetBaseUrl didn't return expected url")
	}
}
