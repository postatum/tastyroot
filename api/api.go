package api

import (
	"net/http"
	"tastyroot/resources"
)

func Register(res resources.ResourceInterface) {
	http.HandleFunc(res.GetUrl(), func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			res.Dehydrate(w, r)
		case "POST":
			res.Hydrate(w, r)
		}
	})
}
