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
	rtr.HandleFunc("/account", RequireCookiesAuthentication(handleGetAccount)).Methods("GET")
	rtr.HandleFunc("/account", handlePostAccount).Methods("POST")
	rtr.HandleFunc("/logout", RequireCookiesAuthentication(handleGetLogout)).Methods("GET")

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

func restAPIrouter() *mux.Router {
	rtr := mux.NewRouter().StrictSlash(true)
	rtr.HandleFunc("/isAuthenticated/{login}/{token}", handleGetIsAuthenticated).Methods("GET")
	rtr.HandleFunc("/prolongAuth/{login}/{token}", handleGetProlongAuth).Methods("GET")
	rtr.HandleFunc("/getAccount/{login}", handleGetAccountJSON).Methods("GET")
	logAndRelay := LogRequestsNoCookies(rtr.ServeHTTP)
	logger := mux.NewRouter()
	logger.PathPrefix("/").HandlerFunc(logAndRelay)
	return logger
}
