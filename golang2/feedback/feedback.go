package feedback

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

func open() *sql.DB {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres "+
		"password=834295 dbname=postgres sslmode=disable")
	checkErr(err)
	return db
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var db *sql.DB

func init() {
	db = open()
	defer db.Close()
	fmt.Println("Successfully connected!")
}

func cors(w http.ResponseWriter) {
	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}

func Addfeedback(w http.ResponseWriter, r *http.Request) {

	cors(w)
	r.ParseForm()
	if r.Method == "POST" {

		result, err := ioutil.ReadAll(r.Body)
		checkErr(err)
		var user map[string]interface{}
		json.Unmarshal(result, &user)
		tuserid := user["userid"]
		tfeedbacktype := user["feedbacktype"]
		tfeedbackcontent := user["feedbackcontent"]
		//fmt.Println(tuserid)
		//fmt.Println(tfeedbacktype)
		//fmt.Println(tfeedbackcontent)
		stmt, err := db.Prepare("insert into feedback(userid,feedbacktype,feedbackcontent,feedbackreplie) values($1,$2,$3,$4)")
		checkErr(err)

		_, err = stmt.Exec(tuserid, tfeedbacktype, tfeedbackcontent, "无")
		checkErr(err)
		//fmt.Println("insert into feedback success")

	}

}

func Userfeedback(w http.ResponseWriter, r *http.Request) {

	cors(w)
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method == "POST" {
		//fmt.Println("123")
		result, err := ioutil.ReadAll(r.Body)
		checkErr(err)
		var user map[string]interface{}
		json.Unmarshal(result, &user)
		tuserid := user["userid"]
		//fmt.Println(tuserid)

		//index := 0 //记录有多少个账户
		rows, err := db.Query("select * from feedback order by feedbackid desc")
		checkErr(err)
		mapInstances := []map[string]interface{}{}
		for rows.Next() {
			var userid string
			var feedbackid string
			var feedbackcontent string
			var feedbacktype string
			var feedbacktime string
			var feedbackreplie string

			err = rows.Scan(&feedbackid, &userid, &feedbackcontent, &feedbacktype, &feedbacktime, &feedbackreplie)
			checkErr(err)

			if tuserid == userid {

				//fmt.Println(feedbacktime)
				instance := map[string]interface{}{"feedbackcontent": feedbackcontent, "feedbacktype": feedbacktype, "feedbacktime": feedbacktime, "feedbackreplie": feedbackreplie}
				mapInstances = append(mapInstances, instance)
				//index++
			}

		}
		allmen := map[string]interface{}{"result": mapInstances}
		jsonStr, err := json.Marshal(allmen)
		checkErr(err)
		w.Write(jsonStr)
		/*"result": */ /*, "index": index*/
	}

}
