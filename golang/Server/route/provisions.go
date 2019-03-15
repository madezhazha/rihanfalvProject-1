package route

import (
	"fmt"
    "log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	//"database/sql"
	_ "github.com/lib/pq"
	"../psql"
)


var Country  string
const (                      //数据库登入信息
    host     = "localhost"
    port     =  5432
    user     = "Dong"
    password = "87257745"
    dbname   = "law"
)

var posttype string
var postlabel string
var posttitle string

type country struct {
    Country        string    `json:"Country"`
}

type legal struct {
    Legalid           int    `json:"legalid"`
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
    Legalcontent       string    `json:"legalcontent"`
    Legallabel        string    `json:"legallabel"`
}
/*
func main() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err = sql.Open("postgres", psqlInfo)
    checkErr(err)
    log.Println("Successful connection to database!")
    log.Println("服务器已开启")
    putincount()
}
*/
func head(w http.ResponseWriter){
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")             //返回数据格式是json
}

func checkErr(err error) {   //报错
    if err != nil {
        log.Println("出错啦!")
        panic(err)
    }
}

func Nowcountry(w http.ResponseWriter, r *http.Request) { 
	head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &country{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    fmt.Println("当前状态：", su.Country)
    Country = su.Country
}

func Typeget(w http.ResponseWriter, r *http.Request) {    //输出信息
    head(w)
    log.Println("开始搜索信息...")
    var types []psql.Legaltype
    if Country=="Japan"{
		types = psql.Typesql()
	}else if Country=="Korea"{
		types = psql.KTypesql()
    }
   data,_:=json.Marshal(types) 
	w.Write(data)
}

func Titlepost(w http.ResponseWriter, r *http.Request) { 
	head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    //if su.Legaltype!=""{
    fmt.Println("客户端访问：")
    fmt.Println("\tlegaltype:", su.Legaltype)
    posttype = su.Legaltype
}

func Titleget(w http.ResponseWriter, r *http.Request) { 
	head(w)
    var titles []psql.Legaltitle
	if Country=="Japan"{
		titles = psql.Titlesql(posttype)
	}
	 if Country=="Korea"{
	 	titles = psql.KTitlesql(posttype)
	}
	data,_:=json.Marshal(titles) 
	w.Write(data)
    log.Println(titles)
}

func Labelpost(w http.ResponseWriter, r *http.Request) { 
    head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    fmt.Println("客户端访问：")
    fmt.Println("\tlegallabel:", su.Legallabel)
    postlabel = su.Legallabel
    
}

func Labelget(w http.ResponseWriter, r *http.Request) { 
    head(w)
	var label []psql.Legaltype
    if Country=="Japan"{
        label=psql.Labelsql(postlabel)
    }else if Country=="Korea"{
        label=psql.KLabelsql(postlabel)
    }
	data,_:=json.Marshal(label) 
    w.Write(data)
    log.Println(label)
}

func Contentpost (w http.ResponseWriter, r *http.Request) { 
    head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    if su.Legaltitle!=""{
        fmt.Println("客户端访问：")
        fmt.Println("\tlegaltitle:", su.Legaltitle)
        posttitle = su.Legaltitle
    }
}

func Contentget (w http.ResponseWriter, r *http.Request) { 
    head(w)
	var content []byte
    
    if Country=="Japan"{
        content=psql.Contentsql(posttitle)
    }else if Country=="Korea"{
        content=psql.KContentsql(posttitle)
    }
    fmt.Fprintf(w,string(content))
}