package main

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	cmt "github.com/kacpekwasny/commontools"
)

// json2map
// and
// checkKeys
// are
// in COMMONTOOLS

// Shortcut for usual response
func Respond(w http.ResponseWriter, r *http.Request, msg_title string, more interface{}) {
	//	msg := GetMessage(msg_title, r)
	err_code := GetErrCode(msg_title)
	w.Header().Add("Content-Type", "application/json")
	resp := map[string]interface{}{
		"err_code": err_code,
		//		"msg":      msg,
	}
	if more != nil {
		resp["more"] = more
	}
	cmt.Pc(json.NewEncoder(w).Encode(resp))
}

func makeMap(key_val_pairs ...string) map[string]string {
	l := len(key_val_pairs)
	if l%2 != 0 {
		panic(ErrInvalidInput)
	}
	m := map[string]string{}
	for i := 0; i <= l-2; i += 2 {
		m[key_val_pairs[i]] = key_val_pairs[i+1]
	}
	return m
}

// ErrPwned
func PasswordPwned(pass string) error {
	sha := crypto.SHA1.New()
	sha.Write([]byte(pass))
	hash := strings.ToUpper(fmt.Sprintf("%x", sha.Sum(nil)))
	hash5 := hash[:5]
	hashend := hash[5:]
	r, err := http.Get("https://api.pwnedpasswords.com/range/" + hash5)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	s := string(body)
	s = strings.ReplaceAll(s, "\r", "")
	ls := strings.Split(s, "\n")
	for i, str := range ls {
		ls[i] = strings.Split(str, ":")[0]
		if ls[i] == hashend {
			return ErrPwned
		}
	}
	return nil
}

// return string of invalid charachters
func LoginValidChars(login string) string {
	const valid = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_ "
	var invalid = ""
	for _, ch := range login {
		chs := string(ch)
		if !strings.Contains(valid, chs) {
			invalid += fmt.Sprintf("'%v', ", chs)
		}
	}
	return invalid
}

func UpdateLastLoginByRequest(r *http.Request) {
	clog, err := r.Cookie("login")
	if err != nil {
		return
	}
	ctok, err := r.Cookie("token")
	if err != nil {
		return
	}
	login, token := clog.Value, ctok.Value
	if M.IsAuthenticated(login, token) == nil {
		M.UpdateLastLogin2Now(login)
	}
}
