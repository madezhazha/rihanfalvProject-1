package head

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkEmailPassword(o output) (input, string) {
	var password string
	var imageUrl string
	var in input
	row, err := db.Query("select password,image from users where email=$1", o.Email)
	defer row.Close()

	check(err)
	if row.Next() {
		err = row.Scan(&password, &imageUrl)
		check(err)
		fmt.Println(password, "2")

		if password == o.Password {
			//放回登陆成功
			in.IfLogin = true
			in.Tip = "登陆成功"
		} else {
			//返回登陆密码错误
			in.IfLogin = false
			in.Tip = "密码错误"
		}
	} else {
		//返回账号不存在
		in.IfLogin = false
		in.Tip = "Email不存在"
	}
	return in, imageUrl
} //登录-查错-

func Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		h(w, r)
	}
} //跨域是函数头

func Hello(w http.ResponseWriter, r *http.Request) {
	var in input
	var imageUrl string

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	check(err)
	fmt.Println(string(body), "******")
	var u output
	err = json.Unmarshal(body, &u)
	in, imageUrl = checkEmailPassword(u)

	if in.IfLogin {
		img, err := ioutil.ReadFile(imageUrl)
		if err != nil {
			fmt.Println("read file error")
		}
		in.Image = img
	}
	data, _ := json.Marshal(in)
	w.Write(data)
} //登录。。。
