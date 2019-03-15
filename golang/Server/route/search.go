
package route

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	_ "github.com/lib/pq"
	"../psql"

)
var text map[string]interface{}
var class[] string
var readkey interface{}
var readclass interface{}
var readcountry interface{}
var readorder interface{}

func Readcontent(w http.ResponseWriter, r *http.Request) { //获取json
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json") //允许跨域
	r.ParseForm()                                      //解析参数，默认是不会解析的
	gettext, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(gettext), &text)
} //包含跨域代码，并获取内容

func M_Search(w http.ResponseWriter, r *http.Request) {
		Readcontent(w,r)
		readkey = text["KeyWord"]//获取搜索内容
		readclass = text["Classify"]//获取搜索内容
		readcountry = text["Nowcountry"]//获取当前国家
		readorder = text["Order"]//获取排序方式
		if readkey==nil{
			return
		}		
		if readclass==nil{
			return
		}	
		if readcountry==nil{
			return
		}
		if readorder==nil{
			return
		}
		fmt.Println("——————进行搜索——————")
		fmt.Println("搜索表:",readclass)
		fmt.Println("关键词:",readkey)
		fmt.Println("搜索方式:",readorder)
		var searchlist []psql.Searchbox
		searchlist=psql.Getclass(readkey.(string),readcountry.(string),readclass.(string),readorder.(string),searchlist)
		psql.Scoreofsearch(searchlist,readkey.(string))//计算搜索的相关度
		psql.SelectSort(searchlist)//按相关度排序
		str, err := json.Marshal(searchlist)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
			}
			if len(searchlist)==0{
				fmt.Fprintf(w,"null")
				fmt.Print("未搜索出数据")
			}
			fmt.Fprintf(w,string(str))
		}