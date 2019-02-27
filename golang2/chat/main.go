package main

import (
	"ChatServe/serve"
	"ChatServe/utils"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("服务器已运行")

	http.HandleFunc("/thread/list", serve.ListThread)
	http.HandleFunc("/thread/post", serve.ListPost)
	http.HandleFunc("/thread/reply", serve.CreatePost)
	http.HandleFunc("/thread/search", serve.Search)
	http.HandleFunc("/thread/collect", serve.Collect)
	http.HandleFunc("/thread/cancel", serve.Cancel)
	// http.HandleFunc("/list", data.Test)

	//指明监听端口
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		utils.Logger.SetPrefix("ERROR ")
		utils.Logger.Println(err)
		return
	}

	// data.SetText("这是此帖子的第二个回帖\n，我们正在测试回帖的文本是否正常工作。\n请无视此回帖")

}
