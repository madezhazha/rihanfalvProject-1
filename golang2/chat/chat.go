package chat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)
var db *sql.DB

//连接
func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// 通过id查个人信息
func ShowUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	r.ParseForm()
	defer r.Body.Close()
	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var postbody Users
	err = json.Unmarshal(s, &postbody)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)
	err1 := db.QueryRow(" select username,password,email,integral from users where userid=$1", &postbody.Userid).Scan(&postbody.Username, &postbody.Password, &postbody.Email, &postbody.Integral)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
		} else {
			log.Fatal(err1)
		}
	}
	rs, err := json.Marshal(postbody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
}

// 通过id查提问列表
func ShowUserQueList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	r.ParseForm()
	defer r.Body.Close()
	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var postbody Getid
	err = json.Unmarshal(s, &postbody) //获取id
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)
	var usertopics []Topics
	rows, err1 := db.Query(" select * from topics where userid=$1 ", postbody.Userid)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var temp Topics
		err := rows.Scan(&temp.Topicid, &temp.Userid, &temp.Topictitle, &temp.Topiccontent, &temp.Numberofreplies, &temp.Collectionvolume, &temp.Visitvolume, &temp.Japanorkorea)

		if err != nil {
			fmt.Println(err)
			return
		}
		usertopics = append(usertopics, temp)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	rs, err := json.Marshal(usertopics)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
	return
}

// 通过id查回答列表
func ShowUserAnsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	r.ParseForm()
	defer r.Body.Close()
	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	var postbody Getid
	err = json.Unmarshal(s, &postbody) //获取id
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)
	var userreplies []Replies
	rows, err1 := db.Query(" select * from replies where userid=$1 ", postbody.Userid)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var temp Replies
		err := rows.Scan(&temp.Replieid,
			&temp.Userid, &temp.Topicid,
			&temp.Replycontent, &temp.Floor)
		if err != nil {
			fmt.Println(err)
			return
		}
		userreplies = append(userreplies, temp)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	rs, err := json.Marshal(userreplies)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, "%s", rs)
}

// 提问
func AddTopics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	r.ParseForm()
	defer r.Body.Close()
	s, err := ioutil.ReadAll(r.Body) //把  body 内容读入字符串 s
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var postbody Topics
	err = json.Unmarshal(s, &postbody)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Test right :", postbody.Userid)
	stmt, err := db.Prepare("insert into topics(userid,topictitle,topiccontent,numberofreplies,collectionvolume,visitvolume,japanorkorea) values($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(postbody.Userid, postbody.Topictitle, postbody.Topiccontent, postbody.Numberofreplies, postbody.Collectionvolume, postbody.Visitvolume, postbody.Japanorkorea)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into user_tbl success")
	}
}

// func main() {
// 	db = connectDB()
// 	http.HandleFunc("/showuserinfo", ShowUserInfo)       // 个人信息
// 	http.HandleFunc("/showuserquelist", ShowUserQueList) // 个人提问列表
// 	http.HandleFunc("/showuseranslist", ShowUserAnsList) // 个人回答列表
// 	http.HandleFunc("/addtopics", AddTopics)             // 添加帖子
// 	err := http.ListenAndServe(":4000", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
