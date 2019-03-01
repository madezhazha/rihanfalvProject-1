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
	//我自己添加的案例分析
	mux.HandleFunc("/alldata",route.Displayhomeall)
	mux.HandleFunc("/displaytxt",route.Displaytxt)
	// 个人主页
	mux.HandleFunc("/get", route.Get)
	mux.HandleFunc("/post", route.Post)

	fmt.Println("Web:7080启动成功")
	err := http.ListenAndServe(":7080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
