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

	fmt.Println("Web:8080启动成功")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
