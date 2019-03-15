package psql

import (
	"fmt"
    "log"
	"encoding/json"	
	//"database/sql"
    _ "github.com/lib/pq"
)
//var db *sql.DB

type legal struct {
    Legalid           int    `json:"legalid"`
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
    Legalcontent       string    `json:"legalcontent"`
    Legallabel        string    `json:"legallabel"`
    Japanorkorea      string    `json:"japanorkorea"`
}

type Legaltype struct {
    Legaltype         string    `json:"legaltype"`
}
type Legaltitle struct {  
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
}
type Legallabel struct {  
    Legallabel        string    `json:"legallabel"`
}
type Legalcontent struct {
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
    Legalcontent       string    `json:"legalcontent"`
}

// func checkErr(err error) {   //报错
//     if err != nil {
//         log.Println("出错啦!")
//         panic(err)
//     }
// }

func Typesql()[]Legaltype{               //从数据库中获取法律总标题
    var Types []Legaltype
    rows, err := db.Query("SELECT distinct legaltype FROM japanlegal;") 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
		Types=append(Types,types)
	}
    rows.Close()
    return Types
}

func KTypesql()[]Legaltype{               //从数据库中获取法律总标题
    var Types []Legaltype
    rows, err := db.Query("SELECT distinct legaltype FROM korealegal;") 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
		Types=append(Types,types)
	}
    rows.Close()
    return Types
}

func Titlesql(legaltype string)[]Legaltitle{               //从数据库中获取法律小标题
    fmt.Println("开始搜索数据库")
    var Titles  []Legaltitle
    rows, err := db.Query("SELECT legaltitle,legaltype FROM japanlegal WHERE legaltype=$1;" ,legaltype) 
    checkErr(err)
    for rows.Next(){
        var titles Legaltitle
		err = rows.Scan(&titles.Legaltitle,&titles.Legaltype)
        checkErr(err)
        if titles.Legaltitle!=""{
		Titles=append(Titles,titles)}
	}
    rows.Close()
    return Titles
}

func KTitlesql(legaltype string)[]Legaltitle{               //从数据库中获取法律小标题
    fmt.Println("开始搜索数据库")
    var Titles  []Legaltitle
    rows, err := db.Query("SELECT legaltitle,legaltype FROM korealegal WHERE legaltype=$1;" ,legaltype) 
    checkErr(err)
    for rows.Next(){
        var titles Legaltitle
		err = rows.Scan(&titles.Legaltitle,&titles.Legaltype)
        checkErr(err)
        if titles.Legaltitle!=""{
		Titles=append(Titles,titles)}
    }
    rows.Close()
    return Titles
}

func Labelsql(legallabel string)[]Legaltype{               //从数据库中获取标签分类
    fmt.Println("开始搜索数据库")
    log.Println(legallabel)
    rows, err := db.Query("SELECT  distinct legaltype FROM japanlegal WHERE legallabel=$1;",legallabel) 
    var Label []Legaltype
    checkErr(err)
    for rows.Next(){
		var label Legaltype
		err = rows.Scan(&label.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
		Label=append(Label,label)
	}
    rows.Close()
    return Label
}

func KLabelsql(legallabel string)[]Legaltype{               //从数据库中获取标签分类
    fmt.Println("开始搜索数据库")
    log.Println(legallabel)
    rows, err := db.Query("SELECT  distinct legaltype FROM korealegal WHERE legallabel=$1;",legallabel) 
    var Label []Legaltype
    checkErr(err)
    for rows.Next(){
		var label Legaltype
		err = rows.Scan(&label.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
		Label=append(Label,label)
	}
    rows.Close()
    return Label
}

func Contentsql (legaltitle string)[]byte{               //从数据库中获取条文正文
    rows, err := db.Query("SELECT distinct legalcontent,legaltype,legaltitle FROM japanlegal WHERE legaltitle = $1;" ,legaltitle) 
    checkErr(err) 
    p := &legal{}
    for rows.Next() {
        var legalcontent string
        var legaltype    string
        var legaltitle   string
        err = rows.Scan(&legalcontent,&legaltype,&legaltitle)
        checkErr(err)
        
        p.Legalcontent = legalcontent
        p.Legaltype = legaltype
        p.Legaltitle = legaltitle
    }
    data, err := json.Marshal(p)
    checkErr(err)
    rows.Close()
    return data
}

func KContentsql (legaltitle string)[]byte{               //从数据库中获取条文正文
    rows, err := db.Query("SELECT distinct legalcontent,legaltype,legaltitle FROM korealegal WHERE legaltitle = $1;" ,legaltitle) 
    checkErr(err) 
    p := &legal{}
    for rows.Next() {
        var legalcontent string
        var legaltype    string
        var legaltitle   string
        err = rows.Scan(&legalcontent,&legaltype,&legaltitle)
        checkErr(err)
        
        p.Legalcontent = legalcontent
        p.Legaltype = legaltype
        p.Legaltitle = legaltitle
    }
    data, err := json.Marshal(p)
    checkErr(err)
    rows.Close()
    return data
}