package route

import (
	//"workplace/psql"
	// "fmt"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	html := "<html><head></head><body><img src='getimage?tag=&name=11.jpg'></body></html>"
	w.Write([]byte(html))
	// fmt.Fprint(w, "hello")
}

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}
