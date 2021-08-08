package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	authserv2 "github.com/kacpekwasny/authserv/src2"
)

func main() {
	fmt.Println("DB manager start...")
	_ = M.Start()
	authserv2.CONFIG.MAX_TOKEN_AGE = time.Hour
	fmt.Printf("Listen on port %v \n", CONFIG_MAP["listen port"])
	r := newRouter()
	log.Fatal(http.ListenAndServe(CONFIG_MAP["listen port"], r))
}
