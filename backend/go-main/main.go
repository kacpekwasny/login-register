package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	cmt "github.com/kacpekwasny/commontools"
)

func main() {
	level, err := strconv.Atoi(CONFIG_MAP["LOG LEVEL"])
	cmt.Pc(err)
	M.LOG_LEVEL = level
	// preapre templates
	LoadConfig()
	var prf = CONFIG_MAP["templates prefix"]
	LoadTemplatesFiles(prf, "login.html", "register.html")
	r := newRouter()

	fmt.Println("Login, register and auth server listening...")
	_ = M.Start()
	log.Println(http.ListenAndServe(CONFIG_MAP["listen port"], r))
}
