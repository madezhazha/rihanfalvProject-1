package route
import(
	"encoding/json"
	"fmt"
	//_ "github.com/bmizerany/pq"
	"io/ioutil"
	"net/http"
	"workplace/psql"
	"encoding/base64"
	"os"
)
//获取用户信息
func Get(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	r.ParseForm() //解析url参数，默认是不会解析的
	if  r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		// 获取前端传来的用户id
		var id int
		json.Unmarshal([]byte(result), &id)
		user1 := psql.SelectUser(id)
		//判断头像是系统头像还是用户本地上传的图像(此非长久之计)
		if(len(user1.Image)<16){
			user1.Image = ImgToBase64(user1.Image)	// 通过路径读取图片，并转成base64传给前端
		}
		js, err := json.Marshal(user1)   //将数据编码成json字符串
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
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
		// 头像文件名是id
		id:=user.UserId
		imgname := fmt.Sprintf("./images/%d.jpg" ,id)
		//判断头像是系统头像还是用户本地上传的图像
		if len(user.Image)>100 {
			Base64ToImg(user.Image,id) //保存前端传来的图片
			user.Image = imgname
		}
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
func ImgToBase64(image string) string {
	img, _ := ioutil.ReadFile(image)               //直接读取文件内容，内容是[]byte类型
	str := base64.StdEncoding.EncodeToString(img)		//将[]byte转化成string
	return str
}
func Base64ToImg(base string, id int) {
	imgname := fmt.Sprintf("./images/%d.jpg" ,id)
	str, _ := base64.StdEncoding.DecodeString(string(base)) // 将base64转化[]byte
	f, _ := os.OpenFile(imgname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(str) //向文件中写入[]byte
}