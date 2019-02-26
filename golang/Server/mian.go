package main

import (
	"fmt"
	"log"
	"net/http"

	"./route"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/test", route.Test)

	fmt.Println("Web:9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
