package thesis

import (
	"database/sql"
	"fmt"
	"strings"
)

//数据库指针
var db *sql.DB

//论文结构体
type Article struct{
	ID int;       //论文ID
	Title string; //标题
	Author string; //作者
	Time string;  //时间
	Content string;  //内容
	Length int;  //长度
	Country string;  //国家
	Lable []string;  //标签组上传给前端
	//Iscollected bool; //是否收藏
	Tagstring string; //数据库标签字符串 用于分割合并处理
}


//连接数据库
func connDb() *sql.DB{
	psqlInfo := "port=5432 user=postgres password=z83313420 dbname=Userinfo sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("open error", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("connect error", err)
		return nil
	}

	fmt.Println("successfull connected!")
	return db
}

//获取所有数据库论文数据       //无标签 方法2
func showArticleList(db *sql.DB) []Article{
	var Articles []Article

	row,erro:=db.Query("SELECT * FROM japanthesis")
	if erro!=nil{
		fmt.Println("Query show ",erro)
	}

	for row.Next(){
		var article Article
		erro = row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Country,&article.Tagstring)
		if erro != nil {
			fmt.Println("showscan error:",erro)
		}
		Articles=append(Articles,article)
	}
	defer row.Close();

	//给论文贴上标签
	for i:=0;i<len(Articles);i++{
		Articles[i].Lable=strings.Split(Articles[i].Tagstring,"/")
	}

	fmt.Println("show over")

	return Articles
}

//分段每次10篇获取论文  page从1开始传入  日本
func getJapanArticlesPart(db *sql.DB,currentpage int) []Article{
	var Articles []Article
	//stmt,erro:=db.Prepare("SELECT * FROM \"Thesis\" LIMIT $1 Where \"JapanOrKorea\"=$2 ")
	//if erro!=nil{
	//	fmt.Println("Query show ",erro)
	//}{
	//fmt.Println(mes)
	rows,erro:=db.Query("SELECT * FROM japanthesis  LIMIT 10 offset $1 ",currentpage)
	if erro!=nil{
		fmt.Println("Query show ",erro)
	}

	for rows.Next(){
		var article Article
		erro = rows.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Country,&article.Tagstring)
		if erro != nil {
			fmt.Println("showscan error:",erro)
		}
		Articles=append(Articles,article)
	}
	defer rows.Close();

	//给论文贴上标签
	for i:=0;i<len(Articles);i++{
		Articles[i].Lable=strings.Split(Articles[i].Tagstring,"/")
	}

	fmt.Println("show over")

	return Articles
}

//分段每次10篇获取论文  page从1开始传入  韩国
func getKoreaArticlesPart(db *sql.DB,currentpage int) []Article{
	var Articles []Article
	//stmt,erro:=db.Prepare("SELECT * FROM \"Thesis\" LIMIT $1 Where \"JapanOrKorea\"=$2 ")
	//if erro!=nil{
	//	fmt.Println("Query show ",erro)
	//}{
	//fmt.Println(mes)
	rows,erro:=db.Query("SELECT * FROM koreathesis  LIMIT 10 offset $1 ",currentpage)
	if erro!=nil{
		fmt.Println("Query show ",erro)
	}

	for rows.Next(){
		var article Article
		erro = rows.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Country,&article.Tagstring)
		if erro != nil {
			fmt.Println("showscan error:",erro)
		}
		Articles=append(Articles,article)
	}
	defer rows.Close();

	//给论文贴上标签
	for i:=0;i<len(Articles);i++{
		Articles[i].Lable=strings.Split(Articles[i].Tagstring,"/")
	}

	fmt.Println("show over")

	return Articles
}

//根据id获取文章详情 日本
func getJapanArticle(db *sql.DB,id int) Article{
	var article Article
	stmt,_:=db.Prepare("Select * from japanthesis where thesisId=$1 ")
	row,_:=stmt.Query(id)
	for row.Next(){
		erro:=row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Country,&article.Tagstring)
		if erro!=nil{
			fmt.Println("Scan erro",erro)
		}
	}
	defer row.Close()
	//插入标签
	article.Lable=strings.Split(article.Tagstring,"/")
	return article
}

//根据id获取文章详情 韩国
func getKoreaArticle(db *sql.DB,id int) Article{
	var article Article
	stmt,_:=db.Prepare("Select * from koreathesis where thesisId=$1 ")
	row,_:=stmt.Query(id)
	for row.Next(){
		erro:=row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Country,&article.Tagstring)
		if erro!=nil{
			fmt.Println("Scan erro",erro)
		}
	}
	defer row.Close()
	//插入标签
	article.Lable=strings.Split(article.Tagstring,"/")
	return article
}