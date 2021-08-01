package main

import (
	"net/http"
	"strings"

	cmt "github.com/kacpekwasny/commontools"
)

var defaultLang = "en"
var allowedLangs = []string{"pl", "en"}

func GetLang(r *http.Request) string {
	ck, err := r.Cookie("lang")
	if err != nil {
		return defaultLang
	}
	lang := strings.Split(ck.String(), "=")[1]
	if _, is := cmt.InSlice(lang, allowedLangs); is {
		return lang
	}
	return defaultLang
}
