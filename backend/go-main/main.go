package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	asrv "github.com/kacpekwasny/authserv/src2"
	cmt "github.com/kacpekwasny/commontools"
)

var M asrv.Manager

func main() {
	LoadConfig()
	// create Manager
	var (
		user          = CONFIG_MAP["user"]
		password      = CONFIG_MAP["password"]
		address       = CONFIG_MAP["address"]
		database_name = CONFIG_MAP["database name"]
		table_name    = CONFIG_MAP["table name"]
	)
	var M = asrv.InitManager(user, password, address, 3306, database_name, table_name,
		time.Second/2, time.Second*2, 20, 1000, time.Minute)

	// set log level
	level, err := strconv.Atoi(CONFIG_MAP["LOG LEVEL"])
	cmt.Pc(err)
	M.LOG_LEVEL = level

	// preapre templates
	var prf = CONFIG_MAP["templates prefix"]
	LoadTemplatesFiles(prf, "login.html", "register.html")

	fmt.Println("DB manager start...")
	_ = M.Start()

	fmt.Printf("Listen on port %v \n", CONFIG_MAP["listen port"])
	r := newRouter()
	log.Fatal(http.ListenAndServe(CONFIG_MAP["listen port"], r))
}
