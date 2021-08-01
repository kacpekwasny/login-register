package main

import (
	"fmt"
	"net/http"

	cmt "github.com/kacpekwasny/commontools"

	asrv "github.com/kacpekwasny/authserv/src2"
)

// w http.ResponseWriter, r *http.Request
func handleGetLogin(w http.ResponseWriter, r *http.Request) {
	lang := GetLang(r)
	m := page_map["login_"+lang]
	ExecuteTemplate(w, "login.html", m)
}

func handlePostLogin(w http.ResponseWriter, r *http.Request) {
	m, err := cmt.Json2mapSS(r)
	if err != nil {
		fmt.Fprint(w, "Cannot decode json")
	}
	missing := cmt.CheckKeys(m, "login", "password")
	if len(missing) > 0 {
		fmt.Fprint(w, "Required json keys: login, password")
	}
	err = M.LoginAccount(m["login"], m["password"])
	fmt.Println("M.LoginAccount(", m["login"], "", err)
	switch err {
	case nil:
		// generate token and set cookie to store it
		acc, _ := M.GetAccount(m["login"]) // account has to still be in cache
		Respond(w, r, "ok", makeMap("token", acc.Current_token))
	case asrv.ErrLoginNotFound:
		Respond(w, r, "login_not_found", makeMap("token", "just-so-it-wouldnt-throw-an-error"))
	case asrv.ErrPasswordMissmatch:
		Respond(w, r, "pass_missmatch", makeMap("token", "just-so-it-wouldnt-throw-an-error"))
	default:
		Respond(w, r, "internal_error", makeMap("token", "just-so-it-wouldnt-throw-an-error"))
	}
}

func handleGetRegister(w http.ResponseWriter, r *http.Request) {
	lang := GetLang(r)
	m := page_map["register_"+lang]
	w.Header().Set("Access-Control-Allow-Origin", "api.pwnedpasswords.com")
	ExecuteTemplate(w, "register.html", m)
}

func handlePostRegister(w http.ResponseWriter, r *http.Request) {
	m, err := cmt.Json2mapSS(r)
	if err != nil {
		fmt.Println("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
	}
	missing := cmt.CheckKeys(m, "login", "password")
	if len(missing) > 0 {
		fmt.Println("Required json keys: login, password", m, "internal error")
		Respond(w, r, "internal_error", nil)
	}
	login := m["login"]
	paswd := m["password"]
	err = PasswordPwned(paswd)
	if err != nil {
		if err == ErrPwned {
			fmt.Println("PasswordPwned error")
			Respond(w, r, "paswd_pwned", nil)
			return
		}
		fmt.Println("PasswordPwned internal error:", err)
		Respond(w, r, "internal_error", nil)
		return
	}
	err = M.RegisterAccount(login, paswd)
	fmt.Println("M.RegisterAccount(", login, ", XXXX )", err)
	switch err {
	case nil:
		Respond(w, r, "ok", nil)
	case asrv.ErrLoginInUse:
		Respond(w, r, "login_in_use", nil)
	case asrv.ErrLoginRequirements:
		Respond(w, r, "login_requirements", nil)
	case asrv.ErrPassRequirements:
		Respond(w, r, "pass_requirements", nil)
	default:
		fmt.Println("handlePostRegister", err)
		Respond(w, r, "internal_error", nil)
	}
}
