package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	cmt "github.com/kacpekwasny/commontools"
)

func main() {
	// preapre templates
	LoadConfig()
	fmt.Println(CONFIG_MAP)

	level, err := strconv.Atoi(CONFIG_MAP["LOG LEVEL"])
	cmt.Pc(err)
	M.LOG_LEVEL = level

	var prf = CONFIG_MAP["templates prefix"]
	LoadTemplatesFiles(prf, "login.html", "register.html")
	r := newRouter()
	fmt.Println("DB manager start...")
	_ = M.Start()
	fmt.Printf("Listen on port %v \n", CONFIG_MAP["listen port"])

	log.Fatal(http.ListenAndServe(CONFIG_MAP["listen port"], r))
}
