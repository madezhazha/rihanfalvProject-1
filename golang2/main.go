package main
 
import(
	"net/http"
	"fmt"
	"time"
	"./homepage"
	"./head"
	"./law"
	"./chat"
	"./analysis"
	"./thesis"
)

func main(){
	fmt.Println("The Porgam is running ...")
	mux := http.NewServeMux()
	//================================================================ homepage
	mux.HandleFunc("/get",homepage.GetImages)
	mux.HandleFunc("/gethomepageartical",homepage.GetHomePageArtical)
	mux.HandleFunc("/gethomepagehotnews",homepage.GetHomePageHotnews)
	//================================================================ law  
    mux.HandleFunc("/type",law.Typeget)
    mux.HandleFunc("/title",law.Titlepost)
    mux.HandleFunc("/label",law.Labelpost)
    mux.HandleFunc("/content",law.Contentpost)
	//================================================================ head
	mux.HandleFunc("/", head.Log(head.Hello))
	mux.HandleFunc("/register", head.Register)
	mux.HandleFunc("/sendVerification", head.Email)
	mux.HandleFunc("/CPsendVerification", head.CPsendVerification)
	mux.HandleFunc("/changePassword", head.ChangePassword)
	//================================================================= analysis
	mux.HandleFunc("/showuserinfo", chat.ShowUserInfo)       // 个人信息
	mux.HandleFunc("/showuserquelist", chat.ShowUserQueList) // 个人提问列表
	mux.HandleFunc("/showuseranslist", chat.ShowUserAnsList) // 个人回答列表
	mux.HandleFunc("/addtopics", chat.AddTopics)             // 添加帖子
	//================================================================= thesis
	mux.HandleFunc("/paper",thesis.articleList)
	mux.HandleFunc("/paperweb",thesis.articleDetial)
	mux.HandleFunc("/paper",thesis.isCollectedArticle)
	//================================================================= ......
	server := &http.Server{
		Addr: "0.0.0.0:4400",
		Handler: mux,
		ReadTimeout: time.Duration(10 * int64(time.Second)),
	}
	server.ListenAndServe()
}
