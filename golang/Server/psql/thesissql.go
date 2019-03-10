package psql

import (
	//"database/sql"
	"fmt"
	"strings"
)

//论文结构体
type Article struct{
	ID int;       //论文ID
	Title string; //标题
	Author string; //作者
	Time string;  //时间
	Content string;  //内容
	Length int;  //长度
	//Country string;  //国家
	Lable []string;  //标签组上传给前端
	Tagstring string; //数据库标签字符串 用于分割合并处理
	IsCollected bool; //是否收藏
}

//收藏结构体
type Colle struct{
	Userid int;
	Collectionid int;
	Collectiontype string
	Collectioncontentid int
	Collectiontime string
}



//获取所有数据库论文数据       //无标签 方法2
func ShowArticleList() []Article{
	var Articles []Article

	row,erro:=db.Query("SELECT * FROM japanthesis")
	if erro!=nil{
		fmt.Println("Query show ",erro)
	}

	for row.Next(){
		var article Article
		erro = row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Tagstring)
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

	//fmt.Println("show over")

	return Articles
}

//分段每次10篇获取论文  page从1开始传入  日本
func GetJapanArticlesPart(currentpage int) []Article{
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
		erro = rows.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Tagstring)
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

	//fmt.Println("show over")

	return Articles
}

//分段每次10篇获取论文  page从1开始传入  韩国
func GetKoreaArticlesPart(currentpage int) []Article{
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
		erro = rows.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Tagstring)
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

	//fmt.Println("show over")

	return Articles
}

//根据id获取文章详情 日本
func GetJapanArticle(id int) Article{
	var article Article
	stmt,_:=db.Prepare("Select * from japanthesis where thesisId=$1 ")
	row,_:=stmt.Query(id)
	for row.Next(){
		erro:=row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Tagstring)
		if erro!=nil{
			fmt.Println("xx Scan erro",erro)
		}
	}
	defer row.Close()
	//插入标签
	article.Lable=strings.Split(article.Tagstring,"/")
	return article
}

//根据id获取文章详情 韩国
func GetKoreaArticle(id int) Article{
	var article Article
	stmt,_:=db.Prepare("Select * from koreathesis where thesisId=$1 ")
	row,_:=stmt.Query(id)
	for row.Next(){
		erro:=row.Scan(&article.ID,&article.Title,&article.Author,&article.Time,&article.Content,&article.Length,&article.Tagstring)
		if erro!=nil{
			fmt.Println("Scan erro",erro)
		}
	}
	defer row.Close()
	//插入标签
	article.Lable=strings.Split(article.Tagstring,"/")
	return article
}

//收藏文章
func CollectArticle(userid int,ctype string,articleid int,collectiontime string){
	stmt,erro:=db.Prepare("Insert into collection (userid,collectiontype,collectioncontentid,collectiontime) values($1,$2,$3,$4)")
	if erro!=nil{
		fmt.Println("collect prepare erro",erro)
	}
	res,erro:=stmt.Exec(userid,ctype,articleid,collectiontime)
	if erro!=nil{
		fmt.Println("stmt erro",erro)
	}
	
	res.LastInsertId()
	
	defer stmt.Close()
	//fmt.Println("insert sucess")
}

//取消收藏
func DeleteArticleCollect(userid int,ctype string,articleid int){
	stmt,erro:=db.Prepare("Delete From collection where userid=$1 and collectiontype=$2 and collectioncontentid=$3")
	if erro!=nil{
		fmt.Println("collect prepare erro",erro)
	}
	res,erro:=stmt.Exec(userid,ctype,articleid)
	if erro!=nil{
		fmt.Println("stmt erro",erro)
	}
	
	res.LastInsertId()
	
	defer stmt.Close()
	//fmt.Println("delete sucess")
}

//查询收藏文章
func GetCollectedArticle(userid int,ctype string,articleid int) bool{
	stmt,_:=db.Prepare("Select * from collection where userid=$1 and collectiontype=$2 and collectioncontentid=$3")
	row,_:=stmt.Query(userid,ctype,articleid)
	var collection Colle
	for row.Next(){
		erro:=row.Scan(&collection.Userid,&collection.Collectionid,&collection.Collectiontype,&collection.Collectioncontentid,&collection.Collectiontime)
		if erro!=nil{
			fmt.Println("Scan erro",erro)
		}
	}
	defer row.Close()
	if collection.Collectioncontentid==0{
		return false
	}else{
		return true
	} 
}