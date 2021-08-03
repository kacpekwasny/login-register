package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	asrv "github.com/kacpekwasny/authserv/src2"
	cmt "github.com/kacpekwasny/commontools"
)

var CONFIG_MAP = LoadConfig()
var M *asrv.Manager = InitManager()

func LoadConfig() map[string]string {
	var conf = map[string]string{}
	f, err := os.Open("conf.json")
	cmt.Pc(err)
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	cmt.Pc(err)
	err = json.Unmarshal(bytes, &conf)
	cmt.Pc(err)
	return conf
}

func InitManager() *asrv.Manager {
	// create Manager
	var (
		user          = CONFIG_MAP["user"]
		password      = CONFIG_MAP["password"]
		address       = CONFIG_MAP["address"]
		database_name = CONFIG_MAP["database name"]
		table_name    = CONFIG_MAP["table name"]
	)
	var m = asrv.InitManager(user, password, address, 3306, database_name, table_name,
		time.Second/2, time.Second*2, 20, 1000, time.Minute)

	// set log level
	level, err := strconv.Atoi(CONFIG_MAP["LOG LEVEL"])
	cmt.Pc(err)
	m.LOG_LEVEL = level
	return m
}

func init() {
	// preapre templates
	var prf = CONFIG_MAP["templates prefix"]
	LoadTemplatesFiles(prf, "login.html", "register.html")
	fmt.Println("end init")
}
