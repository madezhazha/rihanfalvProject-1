package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../psql"
	"strconv"
	"time"
)

//用于按信息查询数据库文章列表
type Mes struct{
	CurrentPage int;  //当前文章列表页数最后文章Id
	Country string;  //当前模块日/韩  
}

//按信息查询文章详情
type Detial struct{
	UserID string;
	ArticleID string;
	Country string;
}

//处理收藏信息结构体
type Collection struct{
	UserID string
	ArticleID string
	Country string
	IsCollected bool
}

//显示论文列表上传
func ArticleList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问..域  跨域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var Articles []psql.Article //文章结构体数组
	var mes Mes
	json.Unmarshal(body,&mes)
	if mes.Country==""{
		return
	}

	if mes.Country=="Japan"{
		Articles=psql.GetJapanArticlesPart(mes.CurrentPage-1)
	}else if mes.Country=="Korea"{
		Articles=psql.GetKoreaArticlesPart(mes.CurrentPage-1)
	}

	data,_:=json.Marshal(Articles) 
	w.Write(data)
	fmt.Println("submit Articles sucess")
}

//获取文章详情
func ArticleDetial(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问..域  跨域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var detial Detial
	//var collletion Collection
	json.Unmarshal(body,&detial)  //collletion
	//id传 0 bug
	if detial.Country==""{     //collletion.ArticleID==0
		return
	}
	var article psql.Article
	articleid,_:=strconv.Atoi(detial.ArticleID)
	userid,_:=strconv.Atoi(detial.UserID)
	if detial.Country=="Japan"{
		article=psql.GetJapanArticle(articleid)  //collletion.ArticleID
		//判断收藏
		if userid!=0{
			article.IsCollected=psql.GetCollectedArticle(userid,"japanthesis",articleid)
		}
	}else if detial.Country=="Korea"{
		article=psql.GetKoreaArticle(articleid)
		if userid!=0{
			article.IsCollected=psql.GetCollectedArticle(userid,"koreathesis",articleid)
		}
	}
	fmt.Println(article.IsCollected)

	data,_:=json.Marshal(article)
	w.Write(data)
	fmt.Println("submit article sucess")
	
}

//处理文章收藏和取消收藏
func IsCollectedArticle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问..域  跨域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var collection Collection
	json.Unmarshal(body,&collection)
	if collection.Country==""{
		return
	}
	articleid,_:=strconv.Atoi(collection.ArticleID)
	userid,_:=strconv.Atoi(collection.UserID)
	if collection.IsCollected==true{
		if collection.Country=="Japan"{
			psql.DeleteArticleCollect(userid,"japanthesis",articleid)
		}else if collection.Country=="Korea"{
			psql.DeleteArticleCollect(userid,"koreathesis",articleid)
		}
	}else if collection.IsCollected==false{
		if collection.Country=="Japan"{
			t:=time.Now().Format("2006-01-02")
			psql.CollectArticle(userid,"japanthesis",articleid,t)
		}else if collection.Country=="Korea"{
			t:=time.Now().Format("2006-01-02")
			psql.CollectArticle(userid,"koreathesis",articleid,t)
		}
	}
	var result map[string]int=map[string]int{"Result":1}
	date,_:=json.Marshal(result)
	w.Write(date)

}