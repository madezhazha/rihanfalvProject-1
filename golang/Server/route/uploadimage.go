package route

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)


//save the images that upload from browser
func Uploadfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var tag = r.PostFormValue("tag")
	r.ParseMultipartForm(32 << 20)        //设置内存大小32mb
	files := r.MultipartForm.File["file"] //获取上传的文件组
	length := len(files)
	//judge if the arry is empty
	if length == 0 {
		w.Write([]byte("× upload empty file... <br> "))
		return
	}
	if(length>20){
		w.Write([]byte("× upload files number more than 20 ...<br> "))
		return
	}
	//read the array and save the files
	for i := 0; i < length; i++ {
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
		filepath := "./driver-test/" + tag + files[i].Filename
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
		w.Write([]byte("√"+ files[i].Filename))
	}
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

//images serve (get images by GET method)
func GetImages2(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	if(r.Method != "GET"){
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"] 	//array
	name := vars["name"]
	if( len(tag)!=1 && len(name)!=1 ){
		fmt.Println("GetImages url worng, vars:",vars)
		return
	}
	//find the images and return []byte
	filepath := "./driver-test/"+tag[0]+name[0];
	temp,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	}
	w.Write(temp)
	return
}

//write html file of upload page to browser
func Html(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	filepath := "./uploadimage.html"
	temp,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	}else {
		w.Write(temp)
	}
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
