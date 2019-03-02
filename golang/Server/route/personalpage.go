package route
import(
	"encoding/json"
	"fmt"
	//_ "github.com/bmizerany/pq"
	"io/ioutil"
	"net/http"
	"../psql"
	"encoding/base64"
	"os"
)
//获取用户信息
func Get(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	user1 := psql.SelectUser()
	js, err := json.Marshal(user1)   //将数据编码成json字符串
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
//修改用户信息
func Post(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm() //解析url参数，默认是不会解析的
	if  r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		//此处不要用httpserve.User，因为传到pgsql.go时，pgsql.User和httpserve.User是不同类型那个的
		var user psql.User
		json.Unmarshal([]byte(result), &user)
		psql.UpdateUser(user)
	}
}
//接收，前端上传图片转成base64
func PostBase(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm() //解析url参数，默认是不会解析的
	if  r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		str, _ := base64.StdEncoding.DecodeString(string(result)) // 将base64转化[]byte
		f, _ := os.OpenFile("picture.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(str) //向文件中写入[]byte
	}
}
//读取图片，转成base64发给前端
func GetBase(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm() //解析url参数，默认是不会解析的
	if  r.Method == "GET" {
		img, _ := ioutil.ReadFile("picture.png")               //直接读取文件内容，内容是[]byte类型
		str := base64.StdEncoding.EncodeToString(img)		//将[]byte转化成string
		fmt.Println(str)
		js, err := json.Marshal(str)   //将数据编码成json字符串
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
}

	// w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")             //返回数据格式是json