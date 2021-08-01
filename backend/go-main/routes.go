package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	// logger takes all traffic loggs it, passes it to rtr and the rtr then responds

	rtr := mux.NewRouter().StrictSlash(true)
	rtr.HandleFunc("/login", LogRequests(handleGetLogin)).Methods("GET")
	rtr.HandleFunc("/login", LogRequests(handlePostLogin)).Methods("POST")
	rtr.HandleFunc("/register", LogRequests(handleGetRegister)).Methods("GET")
	rtr.HandleFunc("/register", LogRequests(handlePostRegister)).Methods("POST")

	var dir = CONFIG_MAP["static dir"]
	fs := http.FileServer(http.Dir(dir))
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	logAndRelay := LogRequests(func(w http.ResponseWriter, r *http.Request) {
		rtr.ServeHTTP(w, r)
	})
	logger := mux.NewRouter()
	logger.PathPrefix("/").HandlerFunc(logAndRelay)

	return logger
}
