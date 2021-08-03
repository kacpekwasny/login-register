package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("DB manager start...")
	_ = M.Start()
	fmt.Printf("Listen on port %v \n", CONFIG_MAP["listen port"])
	r := newRouter()
	log.Fatal(http.ListenAndServe(CONFIG_MAP["listen port"], r))
}
