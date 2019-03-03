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
	//head
	mux.HandleFunc("/", route.Log(route.Hello))
	mux.HandleFunc("/register", route.Register)
	mux.HandleFunc("/sendVerification", route.Email)
	mux.HandleFunc("/CPsendVerification", route.CPsendVerification)
	mux.HandleFunc("/changePassword", route.ChangePassword)
	//homepage
	mux.HandleFunc("/homepage/image", route.GetImages)
	mux.HandleFunc("/homepage/articals", route.GetHomePageArtical)
	mux.HandleFunc("/homepage/hotnews", route.GetHomePageHotnews)

	mux.HandleFunc("/addfeedback", route.Addfeedback)
	mux.HandleFunc("/userfeedback", route.Userfeedback)
	//案例分析
	mux.HandleFunc("/alldata", route.Displayhomeall)
	mux.HandleFunc("/displaytxt", route.Displaytxt)
	mux.HandleFunc("/changecollect",route.CollectData)
	mux.HandleFunc("/InitialState",route.InitialState)

	// 讨论区
	mux.HandleFunc("/thread/list", route.ListThread)
	mux.HandleFunc("/thread/post", route.ListPost)
	mux.HandleFunc("/thread/reply", route.CreatePost)
	mux.HandleFunc("/thread/search", route.Search)
	mux.HandleFunc("/thread/collect", route.Collect)
	mux.HandleFunc("/thread/cancel", route.Cancel)

	// 个人主页
	mux.HandleFunc("/get", route.Get)
	mux.HandleFunc("/post", route.Post)

	//论文
	mux.HandleFunc("/paper", route.ArticleList)                   //论文首页
	mux.HandleFunc("/paperweb", route.ArticleDetial)              //论文内容页
	mux.HandleFunc("/paperweb/collect", route.IsCollectedArticle) //处理论文收藏

	fmt.Println("Web:7080启动成功")
	err := http.ListenAndServe(":7080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}