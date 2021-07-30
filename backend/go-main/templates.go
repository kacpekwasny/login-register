package main

import (
	cmt "imports/commontools"
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplatesFiles(prefix string, paths ...string) {
	// attach prefix to dirs
	dirs2 := []string{}
	for _, p := range paths {
		dirs2 = append(dirs2, prefix+p)
	}
	templates = template.Must(template.ParseFiles(dirs2...))

	//	// if folders have no common prefix - leave it empty: ""
	//	for _, dir := range dirs {
	//		files, err := ioutil.ReadDir(prefix + dir)
	//		if err != nil {
	//			log.Println(err.Error())
	//		}
	//		for _, info := range files {
	//			if strings.HasSuffix(info.Name(), ".html") {
	//
	//			}
	//		}
	//	}
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cmt.Pc(templates.ExecuteTemplate(w, tmpl, data))
}
