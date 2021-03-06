package main

import (
	"./src/api"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	ipAddr := ":8088"
	if len(os.Args) > 1 {
		ipAddr = os.Args[1]
	}

	t := time.Now().Local().Format("2006-01-02 15:04:05 -0700")
	fmt.Printf("%s Listen %s\n", t, ipAddr)

	http.HandleFunc("/article", api.GetArticle)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.ListenAndServe(ipAddr, nil)
}
