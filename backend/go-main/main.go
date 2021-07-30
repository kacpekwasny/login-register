package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	M.LOG_LEVEL = 3
	// preapre templates
	LoadConfig()
	var prf = CONFIG_MAP["templates prefix"]
	LoadTemplatesFiles(prf, "login.html", "register.html")
	r := newRouter()

	fmt.Println("Login, register and auth server listening...")
	_ = M.Start()
	//fmt.Println(M.DELETE_ALL_RECORDS_IN_DATABASE())

	log.Println(http.ListenAndServe(":8080", r))
}
