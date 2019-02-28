package main

import (
	"fmt"
	"log"
	"net/http"

	"./psql"
	"./route"
)

func main() {

	psql.TestDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/test", route.Test)

	fmt.Println("Web:8080启动成功")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
