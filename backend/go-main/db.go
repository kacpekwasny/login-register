package main

import (
	asrv "github.com/kacpekwasny/authserv/src2"
)

//var (
//	user          = CONFIG_MAP["user"]
//	password      = CONFIG_MAP["password"]
//	address       = CONFIG_MAP["address"]
//	database_name = CONFIG_MAP["database name"]
//	table_name    = CONFIG_MAP["table name"]
//)
//
//var M = asrv.InitManager(user, password, address, 3306, database_name, table_name,
//	time.Second/2, time.Second*2, 20, 1000, time.Minute)

// var M = asrv.InitManager("authservuser1", "Nyw)5(pjmL", "10.8.0.1", 3306, "authserv", "accounts1", time.Second/2, time.Second*2, 20, 1000, time.Minute)

var M asrv.Manager
