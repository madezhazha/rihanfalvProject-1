package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"workplace/psql"
	"time"
	"strconv"
)

//处理跨域
func Cross(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	return w
}

//错误的处理
type Response struct {
	Data string `json:"data"`
}

func Displayhomeall(w http.ResponseWriter, r *http.Request) {
	//跨域处理
	w = Cross(w)

	//用来接收数据
	var data map[string]interface{}
	var languageType string
	var NumberCasethingString string

	//接收前端发来的请求的请求
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		var info string = "连接出现错误，请刷新页面"
		response := Response{info}
		json, _ := json.Marshal(response)
		w.Write(json)
	}

	//从post请求中的body中获取请求信息
	json.Unmarshal(body, &data)

	if data != nil {
		if data["content"] == "全部" {
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string)    //这说明是可以的
			//这里从服务端拿去数据
			all_data := psql.Getalldata(languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			//发送数据
			w.Write(json)
		}

		//这是刑事案件、民事案件、行政案例、商事、经济案例的拿去代码
		if data["content"] == "刑事案件" || data["content"] == "民事案件" || data["content"] == "行政案例" || data["content"] == "商事、经济案例" {

			getBody := data["content"].(string)
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string)   
			all_data := psql.Getfirstfloor(getBody,languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}

		if data["content"] == "证据" || data["content"] == "正当防卫" || data["content"] == "自首" || data["content"] == "共同犯罪" {

			getBody := data["content"].(string)
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string) 
			all_data := psql.Getreason(getBody,languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "2019" || data["content"] == "2018" || data["content"] == "2017" || data["content"] == "2016" || data["content"] == "2015" {

			getBody := data["content"].(string)
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string) 
			fmt.Println(getBody)

			all_data := psql.Gettime(getBody,languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "一审" || data["content"] == "二审" || data["content"] == "再审" || data["content"] == "执行" {

			getBody := data["content"].(string)
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string) 
			all_data := psql.Getlevel(getBody,languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "刑法诉讼" || data["content"] == "危害公共安全" || data["content"] == "危害国家安全罪" || data["content"] == "贪贿罪" || data["content"] == "侵犯财产罪" || data["content"] == "合同权纠纷" ||
			data["content"] == "物权纠纷" || data["content"] == "劳动权纠纷" || data["content"] == "人格权纠纷" || data["content"] == "其他纠纷" || data["content"] == "行政机关自行处理案例" || data["content"] == "行政诉讼案" ||
			data["content"] == "证券" || data["content"] == "期货交易" || data["content"] == "保险" || data["content"] == "破产" || data["content"] == "商事仲裁" {

			getBody := data["content"].(string)
			languageType=data["languageType"].(string)
			NumberCasethingString = data["NumberCasething"].(string) 
			all_data := psql.Getsecondfloor(getBody,languageType,NumberCasethingString)

			if(all_data == "系统出现错误"){
				response := Response{all_data.(string)}
				json, _ := json.Marshal(response)
				w.Write(json)
				return
			}

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
	}
}

func Displaytxt(w http.ResponseWriter, r *http.Request) {
	w = Cross(w)

	//解析从前端发来的数据
	var data map[string]interface{}
	var content string
	var userId string

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		var info string = "连接出现错误"
		response := Response{info}
		json, _ := json.Marshal(response)
		w.Write(json)
		return //结束这个函数的运行
	}

	json.Unmarshal(body, &data)

	if data != nil {
		content = data["content"].(string) //这个是标题的内容

		if(data["userid"]!=nil){
			userId = data["userid"].(string)
			fmt.Println(userId)
			all_data:=psql.Gettext_userid(content,userId)
			json, _ := json.Marshal(all_data)
			w.Write(json)
		}else{
			fmt.Println("没有")
			all_data:= psql.Gettext_nouserid(content)
			json, _ := json.Marshal(all_data)
			w.Write(json)
		}
	}
}

//收藏
func CollectData(w http.ResponseWriter,r *http.Request){
	w=Cross(w)

	var Data map[string]interface{}
	var content string
	var insturction string
	var titleId string
	var languageType string
	var userId string
	body,err:=ioutil.ReadAll(r.Body)
	t:=time.Now().Format("2006-01-02")

	if err!=nil{
		fmt.Println(err)
		var info string ="系统出现错误"
		response:=Response{info}
		json,_:=json.Marshal(response)
		w.Write(json)
		return
	}

	json.Unmarshal(body,&Data)
	
	if Data!=nil{
		content = Data["title"].(string)
		insturction = Data["data"].(string)
		titleId = Data["titleId"].(string)  //字符串
		titleID,_:= strconv.Atoi(titleId)     //将字符串转化为整形
		languageType = Data["type"].(string)
		userId = Data["userid"].(string)
		userID,_:= strconv.Atoi(userId)
		
		data:=psql.Implement(content,insturction,titleID,languageType,userID,t)
		if(data=="系统出现错误"){
			response:=Response{data}
			json,_:=json.Marshal(response)
			w.Write(json)
			return
		}
		//给了时间
	}
}


//收藏的初始状态
func InitialState(w http.ResponseWriter,r *http.Request){
	w=Cross(w)

	var Data map[string]interface{}
	var content string
	var titleId string
	var languageType string
	var userId string
	body,err:=ioutil.ReadAll(r.Body)
	

	if err!=nil{
		fmt.Println(err)
		var info string ="系统出现错误"
		response:=Response{info}
		json,_:=json.Marshal(response)
		w.Write(json)
		return
	}

	json.Unmarshal(body,&Data)

	if Data!=nil{
		content = Data["title"].(string)
		titleId = Data["titleId"].(string)
		titleID,_:=strconv.Atoi(titleId)
		languageType = Data["type"].(string)
		userId = Data["userid"].(string)
		userID,_:= strconv.Atoi(userId)
		
		data:=psql.Statecollect(content,titleID,languageType,userID)
		if(data=="系统出现错误"){
			response:=Response{data}
			json,_:=json.Marshal(response)
			w.Write(json)
			return
		}
		response:=Response{data}

		json,_:=json.Marshal(response)
		w.Write(json)
	}
}


//付款
func Payment(w http.ResponseWriter,r *http.Request){
	w = Cross(w)

	var Data map[string]interface{}
	var titleId string
	var userId string
	var integral string

	//查看从前端发来的数据
	body,err:=ioutil.ReadAll(r.Body)

	if err!=nil{
		fmt.Println(err)
		var info string ="连接出现错误"
		response:=Response{info}
		json,_:=json.Marshal(response)
		w.Write(json)
		return
	}
	json.Unmarshal(body,&Data)
	
	if Data!=nil{
		titleId = Data["titleid"].(string)
		userId = Data["userid"].(string)
		integral = Data["integral"].(string)
		data:= psql.Pay(titleId,userId,integral)

		if(data=="系统出现错误"){
			response:=Response{data}
			json,_:=json.Marshal(response)
			w.Write(json)
			return
		}
		response:=Response{data}
		json,_:=json.Marshal(response)
		w.Write(json)
	}

}

//支付宝充值
func Recharge(w http.ResponseWriter,r *http.Request){
	w = Cross(w)

	var Data map[string]interface{}
	var userId string
	var integral string //用户想要充值的积分
	var allintegral string  //用户的所有的积分
	var num_intergral int
	var num_allintergral int

	body,err:=ioutil.ReadAll(r.Body)

	if err!=nil{
		fmt.Println(err)
		var info string ="连接出现错误"
		response:=Response{info}
		json,_:=json.Marshal(response)
		w.Write(json)
		return
	}

	json.Unmarshal(body,&Data)

	if Data!=nil{
		userId = Data["userid"].(string)
		integral = Data["integral"].(string)
		allintegral = Data["allintegral"].(string)
		// fmt.Println(userId,integral,allintegral)
		num_intergral,_ = strconv.Atoi(integral)
		num_allintergral,_=strconv.Atoi(allintegral)
		// fmt.Println(num_allintergral,num_intergral)
		num_allintergral = num_allintergral+num_intergral
		data:=psql.Saveintegral(num_allintergral,userId)
		allintegral = psql.GEtintegral(userId)
		fmt.Println(allintegral)
		fmt.Println(data)
		// if(data=="系统出现错误"){
		// 	response:=Response{data}
		// 	json,_:=json.Marshal(response)
		// 	w.Write(json)
		// 	return	
		// }
		response:=Response{allintegral}
		json,_:=json.Marshal(response)
		w.Write(json)
	}
	fmt.Println("充值成功")
}
