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
		M.Log1("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
		return
	}

	missing := cmt.CheckKeys(m, "login", "password")
	if len(missing) > 0 {
		M.Log1("JSON missing keys! Required: 'login', 'password'. The map: %#v", m)
		Respond(w, r, "internal_error", nil)
		return
	}
	err = M.LoginAccount(m["login"], m["password"])
	M.Log2("M.LoginAccount( %v, *** ) -> %v", m["login"], err)
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
		M.Log1("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
		return
	}
	missing := cmt.CheckKeys(m, "login", "password")
	if len(missing) > 0 {
		M.Log1("JSON missing keys! Required: 'login', 'password'. The map: %#v", m)
		Respond(w, r, "internal_error", nil)
		return
	}
	login := m["login"]
	paswd := m["password"]
	if invalid := LoginValidChars(login); len(invalid) > 0 {
		Respond(w, r, "invalid_chars", makeMap("invalid_chars", invalid))
		return
	}
	err = PasswordPwned(paswd)
	if err == ErrPwned || len(paswd) < 7 {
		M.Log3("PasswordPwned: %v", paswd)
		Respond(w, r, "paswd_pwned", nil)
		return
	} else if err != nil {
		M.Log1("During PasswordPwned( *** ) error occured: %v", err)
	}

	err = M.RegisterAccount(login, paswd)
	M.Log1("RegusterAccount( %v, *** ) -> %v", login, err)
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
		Respond(w, r, "internal_error", nil)
	}
}

func handleGetAccount(w http.ResponseWriter, r *http.Request) {
	// redirect '/login' if user not logged in
	// wrapper takes care of authentication
	// return page
	lang := GetLang(r)
	m := page_map["account_"+lang]
	w.Header().Set("Access-Control-Allow-Origin", "api.pwnedpasswords.com")
	ExecuteTemplate(w, "account.html", m)
}

func handlePostAccount(w http.ResponseWriter, r *http.Request) {
	m, err := cmt.Json2mapSS(r)
	if err != nil {
		M.Log1("Json2mapSS internal error")
		Respond(w, r, "internal_error", nil)
		return
	}
	missing := cmt.CheckKeys(m, "action")
	if len(missing) > 0 {
		M.Log1("JSON missing keys! Required: 'login', 'password'. The map: %#v", m)
		Respond(w, r, "internal_error", nil)
		return
	}
	switch m["action"] {
	case "change":
		missing := cmt.CheckKeys(m, "login", "password", "newLogin", "newPassword")
		if len(missing) > 0 {
			M.Log1("JSON missing keys! Required: login, password, newLogin, newPassword. The map: %#v", m)
			Respond(w, r, "internal_error", nil)
			return
		}
		// authenticate
		login, paswd := m["login"], m["password"]
		acc, err := M.GetAccount(login)
		if err != nil {
			Respond(w, r, "login_not_found", nil)
			return
		}
		if !acc.PasswordMatches(paswd) {
			Respond(w, r, "pass_missmatch", nil)
			return
		}
		nlog, npas := m["newLogin"], m["newPassword"]
		M.Log1("login: '%v', nlog: '%v'", login, nlog)
		if len(nlog) > 0 {
			M.Log3("Change login")
			if invalid := LoginValidChars(nlog); len(invalid) > 0 {
				Respond(w, r, "invalid_chars", makeMap("invalid_chars", invalid))
				return
			}
			if len(nlog) < 2 {
				Respond(w, r, "invalid_chars", nil)
				return
			}
			err = M.UpdateLogin(login, nlog)
			M.Log1("UpdateLogin( %v, %v ) -> %v", login, nlog, err)
			switch err {
			case nil:
				login = nlog
				break
			case asrv.ErrLoginInUse:
				Respond(w, r, "login_in_use", nil)
				return
			case asrv.ErrLoginRequirements:
				Respond(w, r, "login_requirements", nil)
				return
			default:
				Respond(w, r, "internal_error", nil)
				return
			}
		}
		if len(npas) > 0 {
			M.Log3("Change password")
			err = PasswordPwned(npas)
			if err == ErrPwned {
				M.Log3("PasswordPwned: %v", paswd)
				Respond(w, r, "paswd_pwned", nil)
				return
			} else if err != nil {
				M.Log1("During PasswordPwned( *** ) error occured: %v", err)
			}
			if len(npas) < 7 {
				Respond(w, r, "pass", nil)
			}
			acc.SetPassHash(npas)
			err = M.UpdatePassHash(login, acc.Pass_hash)
			M.Log1("UpdatePassHash( %v, ***** ) -> %v", login, err)
			switch err {
			case nil:
				break
			case asrv.ErrLoginInUse:
				Respond(w, r, "login_in_use", nil)
				return
			case asrv.ErrPassRequirements:
				Respond(w, r, "pass_requirements", nil)
				return
			default:
				Respond(w, r, "internal_error", nil)
				return
			}
		}
		Respond(w, r, "ok", nil)
		return

	default:
		Respond(w, r, "internal_error", nil)
	}
}

func handleGetLogout(w http.ResponseWriter, r *http.Request) {
	// request allready authenticated by wrapper
	login, _ := r.Cookie("login") // wrapper ensures that 'login' cookie exists
	M.UpdateLoggedIn(login.Value, false)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
