package route

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func Photo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm() //解析参数，默认是不会解析的

	Rphoto, _ := ioutil.ReadAll(r.Body)
	var IPhoto map[string]interface{}
	json.Unmarshal([]byte(Rphoto), &IPhoto)
	//fmt.Printf("%v", IPhoto)

	if IPhoto != nil {
		InterPhoto := IPhoto["photomsg"]
		InterID := IPhoto["UserID"]

		StrPhoto := InterPhoto.(string)
		FloatUserID := InterID.(float64)

		IntUserID := int(FloatUserID)

		UserID := strconv.Itoa(IntUserID)

		fmt.Println("   StrPhoto:")

		re, _ := regexp.Compile("data:image/png;base64,")

		StrPhoto = re.ReplaceAllString(StrPhoto, "")

		var ImageURL = "../images/" + UserID + ".txt"

		ioutil.WriteFile(ImageURL, []byte(StrPhoto), 0667)

		//读取临时文件
		BaseImg, _ := ioutil.ReadFile(ImageURL)

		//解压
		dist, _ := base64.StdEncoding.DecodeString(string(BaseImg))

		var Image = "../images/" + UserID + ".jpg"

		//写入新文件
		f, _ := os.OpenFile(Image, os.O_RDWR|os.O_CREATE, os.ModePerm)

		defer f.Close()
		f.Write(dist)

	}

}
