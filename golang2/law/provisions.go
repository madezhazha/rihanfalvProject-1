package law

import (
	"fmt"
    "log"
	"io/ioutil"
	"encoding/json"
    "strings"
	"net/http"
	"database/sql"
    _ "github.com/lib/pq"
)

type legal struct {
    Legalid           int    `json:"legalid"`
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
    Legalcontent       string    `json:"legalcontent"`
    Legallabel        string    `json:"legallabel"`
    Japanorkorea      string    `json:"japanorkorea"`
}

const (                      //数据库登入信息
    host     = "localhost"
    port     =  5432
    user     = "Dong"
    password = "87257745"
    dbname   = "law"
)

var db *sql.DB
var title    string
var category string
var content  []byte



func checkErr(err error) {   //报错
    if err != nil {
        log.Println("出错啦!")
        panic(err)
    }
}

func head(w http.ResponseWriter){
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")             //返回数据格式是json
}

func Typeget(w http.ResponseWriter, r *http.Request) {    //输出信息
    head(w)
    log.Println("开始搜索信息...")
	typesql(w)
}

func typesql(w http.ResponseWriter){
    rows, err := db.Query("SELECT distinct legaltype FROM legal; ") 
    checkErr(err)
    columns, err := rows.Columns()
    checkErr(err)
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
for i := range values {
        scanArgs[i] = &values[i]
    }
    list := "["
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        checkErr(err)
        row := "{"
        var value string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            columName := strings.ToLower(columns[i])
            cell := fmt.Sprintf(`"%v":"%v"`, columName, value)
            row = row + cell + ","
        }
        row = row[0 : len(row)-1]
        row += "}"
        list = list + row + ","
    }
    list = list[0 : len(list)-1]
    list += "]"
    
    fmt.Fprintf(w,string(list))
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    rows.Close()
}

func Titlepost(w http.ResponseWriter, r *http.Request) { 
    head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    if su.Legaltype!=""{
    fmt.Println("客户端访问：")
    fmt.Println("\tlegaltype:", su.Legaltype)
    Titlesql(su.Legaltype,w)}
    fmt.Fprintf(w,string(title))
    log.Println(title)
}

func Titlesql(legaltype string, w http.ResponseWriter){
    fmt.Println("开始搜索数据库")
    // rows, err := db.Query("SELECT legaltitle FROM legal WHERE legaltype = '"+legaltype+"';") 
    rows, err := db.Query("SELECT legaltitle FROM legal WHERE legaltype=$1;",legaltype) 
    checkErr(err)
    columns, err := rows.Columns()
    checkErr(err)
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
            scanArgs[i] = &values[i]
        }
    list := "["
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        checkErr(err)
        row := "{"
        var value string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            columName := strings.ToLower(columns[i])
            cell := fmt.Sprintf(`"%v":"%v"`, columName, value)
            row = row + cell + ","
        }
        row = row[0 : len(row)-1]
        row += "}"
        list = list + row + ","
    }
    list = list[0 : len(list)-1]
    title = list + "]"
    
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    rows.Close()
}

func Labelpost(w http.ResponseWriter, r *http.Request) { 
    head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    if su.Legallabel!=""{
    fmt.Println("客户端访问：")
    fmt.Println("\tlegallabel:", su.Legallabel)
    Labelsql(su.Legallabel,w)}
    fmt.Fprintf(w,string(category))
    log.Println(category)
}

func Labelsql(legallabel string, w http.ResponseWriter){
    fmt.Println("开始搜索数据库")
    rows, err := db.Query("SELECT  distinct legaltype FROM legal WHERE legallabel=$1;",legallabel) 
    checkErr(err)
    columns, err := rows.Columns()
    checkErr(err)
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
            scanArgs[i] = &values[i]
        }
    list := "["
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        checkErr(err)
        row := "{"
        var value string
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            columName := strings.ToLower(columns[i])
            cell := fmt.Sprintf(`"%v":"%v"`, columName, value)
            row = row + cell + ","
        }
        row = row[0 : len(row)-1]
        row += "}"
        list = list + row + ","
    }
    list = list[0 : len(list)-1]
    category = ""
    if list !="" {
    category = list + "]"
    fmt.Println(list)}
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    rows.Close()
}

func Contentpost(w http.ResponseWriter, r *http.Request) { 
    head(w)
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body) //获取post的数据
    checkErr(err)
    su := &legal{}         //把json转换回来
    json.Unmarshal([]byte(con), &su)
    if su.Legaltitle!=""{
        fmt.Println("客户端访问：")
        fmt.Println("\tlegaltitle:", su.Legaltitle)
        Contentsql(su.Legaltitle,w)}
        fmt.Fprintf(w,string(content))
}

func Contentsql(legaltitle string, w http.ResponseWriter){
    rows, err := db.Query("SELECT distinct legalcontent,legaltype,legaltitle FROM legal WHERE legaltitle = $1;" ,legaltitle) 
    checkErr(err)
    for rows.Next() {
        var legalcontent string
        var legaltype    string
        var legaltitle   string
        err = rows.Scan(&legalcontent,&legaltype,&legaltitle)
        checkErr(err)
        p := &legal{}
        p.Legalcontent = legalcontent
        p.Legaltype = legaltype
        p.Legaltitle = legaltitle
        data, err := json.Marshal(p)
        checkErr(err)
        content = data
    }
    rows.Close()
}



// func main() {
//     var err error
//     psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//     host, port, user, password, dbname)
//     db, err = sql.Open("postgres", psqlInfo)
//     checkErr(err)
//     log.Println("Successful connection to database!")
//     log.Println("服务器已开启")
//     putincount()
// }


// func putincount(){         //向客户端发出信息
//     http.HandleFunc("/type",typeget)
//     http.HandleFunc("/title",titlepost)
//     http.HandleFunc("/label",labelpost)
//     http.HandleFunc("/content",contentpost)
// 	err:= http.ListenAndServe("localhost:9000", nil)
//     if err != nil {
//         fmt.Println("server failed, err:", err)
//     }
// }
