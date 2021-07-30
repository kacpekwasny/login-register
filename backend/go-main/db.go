package main

import (
	"time"

	asrv "github.com/kacpekwasny/authserv/src2"
)

var M = asrv.InitManager("authservuser1", "Nyw)5(pjmL", "10.8.0.1", 3306, "authserv", "accounts1", time.Second/2, time.Second*2, 20, 1000, time.Minute)
