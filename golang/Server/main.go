package main

import (
	"fmt"
	"log"
	"net/http"

	"workplace/psql"
	"workplace/route"
)

func main() {

	psql.TestDB()

	//psql.GetUserTopic(1)

	mux := http.NewServeMux()

	mux.HandleFunc("/test", route.Test)
	//head
	mux.HandleFunc("/", route.Log(route.Hello))
	mux.HandleFunc("/register", route.Register)
	mux.HandleFunc("/sendVerification", route.Email)
	mux.HandleFunc("/CPsendVerification", route.CPsendVerification)
	mux.HandleFunc("/changePassword", route.ChangePassword)
	//homepage
	mux.HandleFunc("/homepage/articals", route.GetHomePageArtical)
	mux.HandleFunc("/homepage/hotnews", route.GetHomePageHotnews)

	mux.HandleFunc("/addfeedback", route.Addfeedback)
	mux.HandleFunc("/userfeedback", route.Userfeedback)
	//案例分析
	mux.HandleFunc("/alldata", route.Displayhomeall)
	mux.HandleFunc("/displaytxt", route.Displaytxt)
	mux.HandleFunc("/changecollect", route.CollectData)
	mux.HandleFunc("/InitialState", route.InitialState)
	mux.HandleFunc("/payment",route.Payment)
	mux.HandleFunc("/recharge",route.Recharge)

	// 讨论区
	mux.HandleFunc("/thread/list", route.ListThread)
	mux.HandleFunc("/thread/post", route.ListPost)
	mux.HandleFunc("/thread/reply", route.CreatePost)
	mux.HandleFunc("/thread/search", route.Search)
	mux.HandleFunc("/thread/collect", route.Collect)
	mux.HandleFunc("/thread/cancel", route.Cancel)
	mux.HandleFunc("/thread/read", route.Read)

	//讨论区 我的问答
	mux.HandleFunc("/showuserinfo", route.ShowUserInfo)       // 个人信息
	mux.HandleFunc("/showuserquelist", route.ShowUserQueList) // 个人提问列表
	mux.HandleFunc("/showuseranslist", route.ShowUserAnsList) // 个人回答列表
	mux.HandleFunc("/addtopics", route.AddTopics)             // 添加帖子
	mux.HandleFunc("/addtopicvisnum", route.AddTopicVisitNumber)	//阅读量+1
	//mux.HandleFunc("/showreplies", route.ShowReplies)		//显示我的主贴的回复信息

	//法律条文
	mux.HandleFunc("/country", route.Nowcountry)
	mux.HandleFunc("/page", route.Pagepost)
    mux.HandleFunc("/pages", route.Pageget)
	mux.HandleFunc("/type", route.Typeget)
	mux.HandleFunc("/title", route.Titlepost)
	mux.HandleFunc("/titles", route.Titleget)
	mux.HandleFunc("/label", route.Labelpost)
	mux.HandleFunc("/labels", route.Labelget)
	mux.HandleFunc("/content", route.Contentpost)
	mux.HandleFunc("/contents", route.Contentget)
	// 个人主页
	mux.HandleFunc("/get", route.Get)
	mux.HandleFunc("/post", route.Post)
	//上传头像
	mux.HandleFunc("/photo", route.Photo)
	//收藏
	mux.HandleFunc("/collectionthesis", route.UserCollectionThesis)
	mux.HandleFunc("/collectioncase", route.UserCollectionCase)
	mux.HandleFunc("/collectiontopic", route.UserCollectionTopic)
	//论文
	mux.HandleFunc("/paper", route.ArticleList)                   //论文首页
	mux.HandleFunc("/paperweb", route.ArticleDetial)              //论文内容页
	mux.HandleFunc("/paperweb/collect", route.IsCollectedArticle) //处理论文收藏
	//搜索
	mux.HandleFunc("/search", route.M_Search)
	//上传，下载图像服务
	mux.HandleFunc("/upload", route.Uploadfiles)
	mux.HandleFunc("/seefiles/path", route.Seefiles)
	mux.HandleFunc("/images", route.GetImages)
	mux.HandleFunc("/uploadimage", route.Html)

	fmt.Println("Web:7080启动成功")
	err := http.ListenAndServe(":7080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
