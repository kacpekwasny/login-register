package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	// logger takes all traffic loggs it, passes it to rtr and the rtr then responds

	rtr := mux.NewRouter().StrictSlash(true)
	rtr.HandleFunc("/login", handleGetLogin).Methods("GET")
	rtr.HandleFunc("/login", handlePostLogin).Methods("POST")
	rtr.HandleFunc("/register", handleGetRegister).Methods("GET")
	rtr.HandleFunc("/register", handlePostRegister).Methods("POST")
	rtr.HandleFunc("/account", handleGetAccount).Methods("GET")
	rtr.HandleFunc("/account", handlePostAccount).Methods("POST")

	dir, ok := CONFIG_MAP["static dir"]
	if ok {
		fs := http.FileServer(http.Dir(dir))
		rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	}

	// All requests are first handled by logger which then relays them to rtr.
	// logger loggsdata from http.Request
	logAndRelay := LogRequests(rtr.ServeHTTP)
	logger := mux.NewRouter()
	logger.PathPrefix("/").HandlerFunc(logAndRelay)
	return logger
}
