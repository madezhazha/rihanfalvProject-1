package route

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//准备工作：uploadimage.html 里将ip地址换成go运行在的主机的地址
//save the images that upload from browser
func Uploadfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//read and check the tag
	var tag = r.PostFormValue("tag")
	tag, err := check_tag(tag)
	if err != nil {
		w.Write([]byte("× tag set unrignt... <br> "))
		return
	}
	//parse and read the file
	r.ParseMultipartForm(10 << 20)        //得到的文件记录最多maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。
	files := r.MultipartForm.File["file"] //MultipartForm can use only after it multipartform.
	length := len(files)
	//judge if the arry is empty
	if length == 0 {
		w.Write([]byte("× upload empty file... <br> "))
		return
	}
	if length > 20 {
		w.Write([]byte("× upload files number more than 20 ...<br> "))
		return
	}
	//read the array and save the files
	for i := 0; i < length; i++ {
		size := files[i].Size
		name := files[i].Filename
		//limit the file not more than 500kb
		if size > 500<<10 {
			w.Write([]byte("× " + name + " Is bigger than 500kb ...<br> "))
			continue
		}
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		//only accept png and jpg format
		ext := strings.ToLower(path.Ext(files[i].Filename))
		if ext != ".jpg" && ext != ".png" {
			w.Write([]byte("× [" + files[i].Filename + "] is not a image file... <br> "))
			continue
		}
		filename := files[i].Filename
		filepath := "./driver-test/" + tag + filename
		//if alerady have a same name file, don't save the new one
		if judge_exist(filepath) {
			w.Write([]byte("× [" + files[i].Filename + "] already exist... <br> "))
			continue
		}
		//save the file in localhost
		os.Mkdir("./driver-test/"+tag, os.ModePerm)
		cur, err := os.Create(filepath)
		defer cur.Close()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(cur, file)
		fmt.Println("Save file :", filepath)
		imgurl := fmt.Sprint("http://123.207.121.2:7080/images?tag=", tag, "&name=", filename)
		scues_msg := filename + "upload scuess!...<br>"
		img_url := fmt.Sprint("<a target=\"_Blank\" href=", imgurl, ">"+imgurl+"</a><br>")
		feedback := scues_msg + "<!>" + img_url
		w.Write([]byte(feedback))
	}
}

//check if tag can be used or not
func check_tag(tag string) (string, error) {
	if tag == "" {
		return "", nil
	}
	if strings.Contains(tag, "/") == true {
		return tag, errors.New("the tag can't be used.")
	}
	return tag + "/", nil
}

//reponse the images alerady have
func Seefiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"]
	w.Write([]byte(filelist("./driver-test/" + tag[0])))
}

//judge if the file is exist in a certain path
func judge_exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//create a html text that descript the part of a certain path
func filelist(path string) (list string) {
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		return "no file in " + path + "...."
	}
	var buffer strings.Builder
	for _, v := range dir_list {
		buffer.WriteString(v.Name() + "<br>")
	}
	list = buffer.String()
	fmt.Println(list)
	return
}

//images serve
func GetImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"] //array
	name := vars["name"]
	if len(tag) != 1 && len(name) != 1 {
		fmt.Println("GetImages url worng, vars:", vars)
		return
	}
	//find the images and return []byte
	filepath := "./driver-test/" + tag[0] + name[0]
	temp, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	}
	w.Write(temp)
	return
}

//write html file of upload page to browser
func Html(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	filepath := "./route/uploadimage.html"
	temp, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	} else {
		w.Write(temp)
	}
}
