package main

import (
	"fmt"
	"net/http"
)

func LogRequests(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		handler(w, r)
	}
}
