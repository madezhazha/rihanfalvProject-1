package analysis

import(
	"database/sql"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"./drivesql"
)

//处理跨域
func Cross(w http.ResponseWriter)http.ResponseWriter{
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json
	return w
}

//打开数据库，可以使用一个函数将他包装
func Opensql(dbname string)*sql.DB{
	db :=drivesql.GetDB(dbname)
	return db
}

//错误的处理
type Response struct{
	Data string `json:"data"`
}


func Displayhomeall(w http.ResponseWriter, r *http.Request){
	//跨域处理
	w = Cross(w)

	//用来接收数据
	var data map[string]interface{}

	db:=Opensql("simplelawcontent")
	fmt.Println(db)

	//接收前端发来的请求的请求
	body,err:= ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println(err)
		var info string="连接出现错误，请刷新页面"
		response := Response{info}
		json, _ := json.Marshal(response)
		w.Write(json)
	}

	//从post请求中的body中获取请求信息
	json.Unmarshal(body,&data)

	if data!=nil{
		if data["content"]=="全部"{
			//这里从服务端拿去数据
			all_data := drivesql.Getalldata(db)
			// fmt.Println(all_data)

			json,_:= json.Marshal(all_data)

			//发送数据
			w.Write(json)
		}
		
		//这是刑事案件、民事案件、行政案例、商事、经济案例的拿去代码
		if (data["content"]=="刑事案件"||data["content"]=="民事案件"||data["content"]=="行政案例"||data["content"]=="商事、经济案例"){

			getBody:=data["content"].(string)
			
			all_data := drivesql.Getfirstfloor(db,getBody)

			json,_:= json.Marshal(all_data)

			w.Write(json)
		}

		if (data["content"]=="证据"||data["content"]=="正当防卫"||data["content"]=="自首"||data["content"]=="共同犯罪"){

			getBody:=data["content"].(string)
			
			all_data := drivesql.Getreason(db,getBody)

			json,_:= json.Marshal(all_data)

			w.Write(json)
		}
		if (data["content"]=="2019"||data["content"]=="2018"||data["content"]=="2017"||data["content"]=="2016"||data["content"]=="2015"){
			
			getBody:=data["content"].(string)

			fmt.Println(getBody)

			all_data := drivesql.Gettime(db,getBody)

			json,_:= json.Marshal(all_data)

			w.Write(json)
		}
		if (data["content"]=="一审"||data["content"]=="二审"||data["content"]=="再审"||data["content"]=="执行"){

			getBody:=data["content"].(string)
			
			all_data := drivesql.Getlevel(db,getBody)

			json,_:= json.Marshal(all_data)

			w.Write(json)
		}
		if (data["content"]=="刑事诉讼"||data["content"]=="危害公共安全"||data["content"]=="危害国家安全罪"||data["content"]=="贪贿罪"||data["content"]=="侵犯财产罪"||data["content"]=="合同权纠纷"||
		data["content"]=="物权纠纷"||data["content"]=="劳动权纠纷"||data["content"]=="人格权纠纷"||data["content"]=="其他纠纷"||data["content"]=="行政机关自行处理案例"||data["content"]=="行政诉讼案"||
		data["content"]=="证券"||data["content"]=="期货交易"||data["content"]=="保险"||data["content"]=="破产"||data["content"]=="商事仲裁"){
			
			getBody:=data["content"].(string)

			all_data := drivesql.Getsecondfloor(db,getBody)

			json,_:= json.Marshal(all_data)

			w.Write(json)
		}
	}
}


func Displaytxt(w http.ResponseWriter, r *http.Request){
	w=Cross(w)

	db:=Opensql("concretelaw")


	//解析从前端发来的数据
	var data map[string]interface{}
	var content string

	body,err:=ioutil.ReadAll(r.Body)

	if err!=nil{
		fmt.Println(err)
		var info string="连接出现错误"
		response:=Response{info}
		json,_:=json.Marshal(response)
		w.Write(json)
		return //结束这个函数的运行
	}

	json.Unmarshal(body,&data)

	if data!=nil{
		content = data["content"].(string)

		//根据这个部分处理从前端发来的请求
		all_data:=drivesql.Gettext(db,content)

		json,_:= json.Marshal(all_data)

		w.Write(json)
	}
}



//测试数据
func Printstring(t string,w string){
	fmt.Println(t,w)
}

