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

	mux.HandleFunc("/addfeedback", route.Addfeedback)
	mux.HandleFunc("/userfeedback", route.Userfeedback)

	fmt.Println("Web:7080启动成功")
	err := http.ListenAndServe(":7080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
