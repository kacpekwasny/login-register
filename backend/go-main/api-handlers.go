package main

import (
	"net/http"

	cmt "github.com/kacpekwasny/commontools"
)

func handleGetIsAuthenticated(w http.ResponseWriter, r *http.Request) {
	m, err := cmt.Json2mapSS(r)
	if err != nil {
		M.Log1("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
	}
	missing := cmt.CheckKeys(m, "login", "token")
	if len(missing) > 0 {
		M.Log1("JSON missing keys! Required: 'login', 'password'. The map: %#v", m)
		Respond(w, r, "internal_error", nil)
		return
	}
	if M.IsAuthenticated(m["login"], m["token"]) != nil {
		Respond(w, r, "unauth", nil)
		return
	}
	Respond(w, r, "ok", nil)
}

func handleGetAccountJSON(w http.ResponseWriter, r *http.Request) {
	m, err := cmt.Json2mapSS(r)
	if err != nil {
		M.Log1("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
	}
	missing := cmt.CheckKeys(m, "login", "token")
	if len(missing) > 0 {
		M.Log1("JSON missing keys! Required: 'login', 'password'. The map: %#v", m)
		Respond(w, r, "internal_error", nil)
		return
	}
	acc, _ := M.GetAccount(m["login"])
	if acc == nil {
		Respond(w, r, "login_not_found", nil)
		return
	}
	Respond(w, r, "ok", acc)
}
