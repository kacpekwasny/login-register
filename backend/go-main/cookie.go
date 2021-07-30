package main

import (
	"fmt"
	cmt "imports/commontools"
	"net/http"
	"strings"
)

var defaultLang = "en"
var allowedLangs = []string{"pl", "en"}

func GetLang(r *http.Request) string {
	ck, err := r.Cookie("lang")
	if err != nil {
		return defaultLang
	}
	lang := strings.Split(ck.String(), "=")[1]
	fmt.Println(lang)
	if _, is := cmt.InSlice(lang, allowedLangs); is {
		return lang
	}
	return defaultLang
}
