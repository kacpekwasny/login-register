package main

import (
	"fmt"
	"net/http"
)

func LogRequests(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Proto, r.Method, r.RequestURI, r.Header.Get("x-forwarded-for"))
		handler(w, r)
	}
}
