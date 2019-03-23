package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../psql"
	_ "github.com/lib/pq"
)
//接收反馈数据
func Addfeedback(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm()
	if r.Method == "POST" {
		result, err := ioutil.ReadAll(r.Body)
		CheckErr(err)
		var user map[string]interface{}
		json.Unmarshal(result, &user)
		tuserid := user["userid"]
		tfeedbacktype := user["feedbacktype"]
		tfeedbackcontent := user["feedbackcontent"]
		//fmt.Println(tuserid)
		//fmt.Println(tfeedbacktype)
		//fmt.Println(tfeedbackcontent)
		psql.Addfeedback(tuserid.(float64), tfeedbacktype.(string), tfeedbackcontent.(string))
		//fmt.Println("insert into feedback success")
	}
}
//接收用户id
func Userfeedback(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method == "POST" {
		//fmt.Println("123")
		result, err := ioutil.ReadAll(r.Body)
		CheckErr(err)
		var user map[string]interface{}
		json.Unmarshal(result, &user)
		tuserid := user["userid"]
		//fmt.Println(tuserid)
		var Feedbacklist []psql.Feedback
		Feedbacklist = psql.Userfeedback(tuserid.(float64))
		res, _ := json.Marshal(Feedbacklist)
		w.Write(res)
	}
}
