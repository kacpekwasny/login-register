package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	authserv2 "github.com/kacpekwasny/authserv/src2"
)

func main() {
	fmt.Println("DB manager start...")
	_ = M.Start()
	authserv2.CONFIG.MAX_TOKEN_AGE = time.Hour
	authserv2.CONFIG.TOKEN_LENGTH = 80
	fmt.Printf("Main router listen on port %v \n", CONFIG_MAP["listen port"])
	fmt.Printf("API router listen on port %v \n", CONFIG_MAP["api listen address and port"])
	r := newRouter()
	rapi := restAPIrouter()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Println(http.ListenAndServe(CONFIG_MAP["listen port"], r))
		fmt.Println("Main router returned.")
	}()
	go func() {
		defer wg.Done()
		log.Println(http.ListenAndServe(CONFIG_MAP["api listen address and port"], rapi))
		fmt.Println("API router returned.")
	}()
	wg.Wait()
}
