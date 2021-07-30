package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	M.LOG_LEVEL = 3
	// preapre templates
	const prf = "/home/kacper/Code/PRODUCTION/web_server/kwasnyy.pl/servers/login-register/templates/"
	LoadTemplatesFiles(prf, "login.html", "register.html")
	r := newRouter()

	fmt.Println("Login, register and auth server listening...")
	_ = M.Start()
	//fmt.Println(M.DELETE_ALL_RECORDS_IN_DATABASE())

	log.Println(http.ListenAndServe(":8080", r))
}
