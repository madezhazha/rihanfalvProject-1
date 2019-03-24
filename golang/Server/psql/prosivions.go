package psql

import (
	"fmt"
    "log"
	"encoding/json"	
	//"database/sql"
    _ "github.com/lib/pq"
)
//var db *sql.DB

type page struct {
    Allpage           int       `json:"allpage"`
}

type legal struct {
    Legalid           int    `json:"legalid"`
    Legaltype         string    `json:"legaltype"`
    Legaltitle        string    `json:"legaltitle"`
    Legalcontent       string    `json:"legalcontent"`
    Legallabel        string    `json:"legallabel"`
    Japanorkorea      string    `json:"japanorkorea"`
}

type Legaltype struct {
    Legalid           int       `json:legalid`
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

func Typesqlp()[]byte{
    var Page page
    var i=0
    rows, err := db.Query("SELECT distinct legaltype FROM japanlegal;") 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        i++
	}
    rows.Close()
    log.Println(i)
    Page.Allpage = i
    data, err := json.Marshal(Page)
    return data
}

func Typesqlpk()[]byte{
    var Page page
    var i=0
    rows, err := db.Query("SELECT distinct legaltype FROM korealegal;") 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        i++
        Page.Allpage = i
	}
    rows.Close()
    log.Println(i)
    data, err := json.Marshal(Page)
    return data
}

func Typesqlpc(legallabel string)[]byte{
    var Page page
    var i=0
    rows, err := db.Query("SELECT  distinct legaltype FROM japanlegal WHERE legallabel=$1;",legallabel) 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        i++
	}
    rows.Close()
    log.Println(i)
    Page.Allpage = i
    data, err := json.Marshal(Page)
    return data
}

func Typesqlpkc(legallabel string)[]byte{
    var Page page
    var i=0
    rows, err := db.Query("SELECT  distinct legaltype FROM korealegal WHERE legallabel=$1;",legallabel) 
    checkErr(err)
    for rows.Next(){
		var types Legaltype
		err = rows.Scan(&types.Legaltype)
		if err != nil {
			fmt.Println("showscan error:",err)
		}
        i++
        Page.Allpage = i
	}
    rows.Close()
    log.Println(i)
    data, err := json.Marshal(Page)
    return data
}

func Titlesqlp(legaltype string)[]byte{
    var Page page
    var i = 0
    rows, err := db.Query("SELECT legaltitle,legaltype FROM japanlegal WHERE legaltype=$1;" ,legaltype) 
    checkErr(err)
    for rows.Next(){
        var titles Legaltitle
		err = rows.Scan(&titles.Legaltitle,&titles.Legaltype)
        checkErr(err)
        i++
	}
    rows.Close()
    log.Println(i)
    Page.Allpage = i
    data, err := json.Marshal(Page)
    return data
}

func Titlesqlpk(legaltype string)[]byte{
    var Page page
    var i = 0
    rows, err := db.Query("SELECT legaltitle,legaltype FROM korealegal WHERE legaltype=$1;" ,legaltype) 
    checkErr(err)
    for rows.Next(){
        var titles Legaltitle
		err = rows.Scan(&titles.Legaltitle,&titles.Legaltype)
        checkErr(err)
        i++
	}
    rows.Close()
    log.Println(i)
    Page.Allpage = i
    data, err := json.Marshal(Page)
    return data
}

func Typesql(page int)[]Legaltype{               //从数据库中获取法律总标题
    var Types []Legaltype
    fmt.Println(page)
    rows, err := db.Query("SELECT distinct legaltype FROM japanlegal limit 20 offset $1;",page) 
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

func KTypesql(page int)[]Legaltype{               //从数据库中获取法律总标题
    var Types []Legaltype
    rows, err := db.Query("SELECT distinct legaltype FROM korealegal limit 10 offset $1;",page) 
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

func Titlesql(legaltype string,page int)[]Legaltitle{               //从数据库中获取法律小标题
    fmt.Println("开始搜索数据库")
    var Titles  []Legaltitle
    rows, err := db.Query("SELECT legaltitle,legaltype FROM japanlegal WHERE legaltype=$1 limit 20 offset $2;" ,legaltype ,page) 
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

func KTitlesql(legaltype string,page int)[]Legaltitle{               //从数据库中获取法律小标题
    fmt.Println("开始搜索数据库")
    var Titles []Legaltitle
    rows, err := db.Query("SELECT legaltitle,legaltype FROM korealegal WHERE legaltype=$1 limit 20 offset $2;" ,legaltype ,page) 
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

func Labelsql(legallabel string,page int)[]Legaltype{               //从数据库中获取标签分类
    fmt.Println("开始搜索数据库")
    log.Println(legallabel)
    rows, err := db.Query("SELECT  distinct legaltype FROM japanlegal WHERE legallabel=$1 limit 1 offset $2;",legallabel,page) 
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

func KLabelsql(legallabel string,page int)[]Legaltype{               //从数据库中获取标签分类
    fmt.Println("开始搜索数据库")
    log.Println(legallabel)
    rows, err := db.Query("SELECT  distinct legaltype FROM korealegal WHERE legallabel=$1 limit 1 offset $2;",legallabel,page) 
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