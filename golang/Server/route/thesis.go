package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"../psql"
	"strconv"
	
)



//用于按信息查询数据库文章列表
type Mes struct{
	CurrentPage int;  //当前文章列表页数最后文章Id
	Country string;  //当前模块日/韩  
}

//按信息查询文章详情
type Detial struct{
	ID string;
	Country string;
}


//显示论文列表上传
func articleList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问..域  跨域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var Articles []psql.Article //文章结构体数组
	var mes Mes
	json.Unmarshal(body,&mes)
	fmt.Println(mes)
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
func articleDetial(w http.ResponseWriter, r *http.Request){
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
	id,_:=strconv.Atoi(detial.ID)
	if detial.Country=="Japan"{
		article=psql.GetJapanArticle(id)  //collletion.ArticleID
	}else if detial.Country=="Korea"{
		article=psql.GetKoreaArticle(id)
	}
	
	//判断是否收藏
	//IsCollected:=getUserCollectiond(psql.DB,collletion)
	//article.Iscollected=IsCollected

	data,_:=json.Marshal(article)
	w.Write(data)
	fmt.Println("submit article sucess")
	
}

//处理文章收藏和取消收藏
func isCollectedArticle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问..域  跨域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	defer r.Body.Close()
	//
}