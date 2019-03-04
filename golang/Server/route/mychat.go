package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"../psql"
	// "github.com/ascoders/alipay"
)

//处理跨域
// func Cross(w http.ResponseWriter) http.ResponseWriter {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
// 	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
// 	w.Header().Set("content-type", "application/json")             //返回数据格式是json
// 	return w
// }

//测试显示时间数据
func Query_test_time(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)

	defer r.Body.Close()

	s, _ := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	var temp psql.Getid
	err := json.Unmarshal(s, &temp)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", temp.Userid)

	var postbody psql.Reply
	postbody=psql.Query_time(temp.Userid)

	rs, err := json.Marshal(postbody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
	return
}

// 通过id查个人信息
func ShowUserInfo(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)
	
	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var postbody psql.Users
	err = json.Unmarshal(s, &postbody)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)

	postbody=psql.GetUserInfo(postbody)
	
	rs, err := json.Marshal(postbody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
}

// 通过id查提问列表
func ShowUserQueList(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)

	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var postbody psql.Getid
	err = json.Unmarshal(s, &postbody) //获取id
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)

	var usertopics []psql.Topics

	usertopics=psql.GetUserTopic(postbody.Userid)
	
	rs, err := json.Marshal(usertopics)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
	return
}

// 通过id查回答列表
func ShowUserAnsList(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)

	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var postbody psql.Getid
	err = json.Unmarshal(s, &postbody) //获取id
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)

	var userreplies []psql.Replies
	userreplies=psql.GetUserReply(postbody)
	
	rs, err := json.Marshal(userreplies)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
}

// 提问
func AddTopics(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)

	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var postbody psql.Topics
	err = json.Unmarshal(s, &postbody)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)

	psql.InsertTopic(postbody)

	rs, err := json.Marshal(postbody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
}