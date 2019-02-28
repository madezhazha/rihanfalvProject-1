package route

import (
	//"../psql"
	"fmt"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func Cors(w http.ResponseWriter) {
	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}
