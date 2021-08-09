package main

import (
	"fmt"
	"net/http"
)

func LogRequests(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("x-forwarded-for"), r.Method, r.RequestURI,
			"\n ⤷", Cookie2Str(r))
		handler(w, r)
	}
}

func LogRequestsNoCookies(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("x-forwarded-for"), r.Method, r.RequestURI)
		handler(w, r)
	}
}

func RequireCookiesAuthentication(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// redirect '/login' if user not logged in
		clog, err := r.Cookie("login")
		if err != nil {
			M.Log2("Request cookie without login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		ctok, err := r.Cookie("token")
		if err != nil {
			M.Log2("Request cookie without token")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		login, token := clog.Value, ctok.Value
		err = M.IsAuthenticated(login, token)
		if err != nil {
			M.Log2("%v is not authenticated, err: %v", login, err)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		M.UpdateLastLogin2Now(login)
		handler(w, r)
	}
}
