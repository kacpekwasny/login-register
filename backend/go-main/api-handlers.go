package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleGetIsAuthenticated(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	if M.IsAuthenticated(m["login"], m["token"]) != nil {
		Respond(w, r, "unauth", nil)
		return
	}
	Respond(w, r, "ok", nil)
}

func handleGetAccountJSON(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	acc, _ := M.GetAccount(m["login"])
	if acc == nil {
		Respond(w, r, "login_not_found", nil)
		return
	}
	Respond(w, r, "ok", acc)
}
