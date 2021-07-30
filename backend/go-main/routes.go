package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/login", LogRequests(handleGetLogin)).Methods("GET")
	r.HandleFunc("/login", LogRequests(handlePostLogin)).Methods("POST")
	r.HandleFunc("/register", LogRequests(handleGetRegister)).Methods("GET")
	r.HandleFunc("/register", LogRequests(handlePostRegister)).Methods("POST")

	// static files
	// const dir = "/media/kacper/Transcend_K1/Home/IT/Code/github.com/kacpekwasny/login-register-page/static/"
	var dir = CONFIG_MAP["static dir"]
	fs := http.FileServer(http.Dir(dir))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}
