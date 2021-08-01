package main

import (
	"net/http"
	"text/template"

	cmt "github.com/kacpekwasny/commontools"
)

var templates *template.Template

func LoadTemplatesFiles(prefix string, paths ...string) {
	// attach prefix to dirs
	dirs2 := []string{}
	for _, p := range paths {
		dirs2 = append(dirs2, prefix+p)
	}
	templates = template.Must(template.ParseFiles(dirs2...))
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cmt.Pc(templates.ExecuteTemplate(w, tmpl, data))
}
