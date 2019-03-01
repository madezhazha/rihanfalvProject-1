package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../psql"
	// "github.com/ascoders/alipay"
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
			//这里从服务端拿去数据
			all_data := psql.Getalldata()
			// fmt.Println(all_data)

			json, _ := json.Marshal(all_data)

			//发送数据
			w.Write(json)
		}

		//这是刑事案件、民事案件、行政案例、商事、经济案例的拿去代码
		if data["content"] == "刑事案件" || data["content"] == "民事案件" || data["content"] == "行政案例" || data["content"] == "商事、经济案例" {

			getBody := data["content"].(string)

			all_data := psql.Getfirstfloor(getBody)

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}

		if data["content"] == "证据" || data["content"] == "正当防卫" || data["content"] == "自首" || data["content"] == "共同犯罪" {

			getBody := data["content"].(string)

			all_data := psql.Getreason(getBody)

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "2019" || data["content"] == "2018" || data["content"] == "2017" || data["content"] == "2016" || data["content"] == "2015" {

			getBody := data["content"].(string)

			fmt.Println(getBody)

			all_data := psql.Gettime(getBody)

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "一审" || data["content"] == "二审" || data["content"] == "再审" || data["content"] == "执行" {

			getBody := data["content"].(string)

			all_data := psql.Getlevel(getBody)

			json, _ := json.Marshal(all_data)

			w.Write(json)
		}
		if data["content"] == "刑事诉讼" || data["content"] == "危害公共安全" || data["content"] == "危害国家安全罪" || data["content"] == "贪贿罪" || data["content"] == "侵犯财产罪" || data["content"] == "合同权纠纷" ||
			data["content"] == "物权纠纷" || data["content"] == "劳动权纠纷" || data["content"] == "人格权纠纷" || data["content"] == "其他纠纷" || data["content"] == "行政机关自行处理案例" || data["content"] == "行政诉讼案" ||
			data["content"] == "证券" || data["content"] == "期货交易" || data["content"] == "保险" || data["content"] == "破产" || data["content"] == "商事仲裁" {

			getBody := data["content"].(string)

			all_data := psql.Getsecondfloor(getBody)

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
		//根据这个部分处理从前端发来的请求
		all_data := psql.Gettext(content) //这个是法官的观点
		json, _ := json.Marshal(all_data)

		w.Write(json)
	}
}

// func Testalipay(w http.ResponseWriter, r *http.Request) {
// 	//这里测试能否实现支付宝支付

// 	w = Cross(w)

// 	Alipay := alipay.Client{
// 		Partner:   "2016092800613779",
// 		Key:       "MIIEpAIBAAKCAQEA7BJPNAZ/M3fbMb7xOQQaAwb/bRTIU8M3LU/QVzCNiSROZpUkKrFeeEPsSD+9Fq0CgZ2FqGnKWl6Pyhri32W/x9e43GY5dpDtd8HZuOwSUHIN89KbKa7qtRx2F831NvARO2HpwYFBtFHcMvbjTCoT+phtx3drILgk3gxoZ45SCaOet77/xBuP5GvlU4N/5Ksprm2CNEV2tXr4j8fIcG34/aWwA54zdFc0+8aytfjHsQ445u4ZkcoUiMrJ1XF9xn8m6el0dzKak9rL8zGcRrsBVYwh6x/Tgsl7uWSKS2scs4fdIubtigOOhi7GSMQY6PVBBwK7b604gD4U8ejQeM/LiwIDAQABAoIBAB4mUHtCgShfnF0EF+bCQVg1opWZ2+PmwLtGwEMCbnEfc+viDFZvhLMzaY/opAurWPCY36Mcr71O6mpTeLozUoUBZhiv7ZLY+5sZ5OabL9OaXXHQagSu1EcBNYvgv/qeOnUtEh+mlG+lLMxNf8YQcbDFmu+orsPlAMBjMUpYgN2+CwkCijjWQ5WoFtpq6yCBg/dIA8/FCe9MJBB52VwR1amLX+9uM/ps3ayPctF9U5HMUV8D2ddqRkq1BVLydO/uEhxMC2rgYFf3r66/nKDgNLHGcdYhDRYdA/Cee2dqQKl0nd+KpnIxyTo1MXl9FH/agOl1eIRi9VYrFVaJGyWqXFECgYEA/jfcHIwi4+jQ9beFayMO6ILayzxPU0RFGRhMwvLSfzmyfyoKWRNni7XzVBtB0Q47CAq/KBotKSwtORB1SWaryOWECUyCxCXszVO1TSbBq69Iyvo4wwCIw46BsJHHSdoIkyr8b7mCUjvh1QDIplEcpZBxqpZeelM1cMD2a10P8X0CgYEA7bnjpSY3iuRPITQR+byDd0gzny4CsACnXUa644sW7Ta8Zcya5+DRj+eG7MxCN5CpIXcnJflp+RherlDjvWKb0qiimSOY7CJxpR3i2bkeyMYJe97kJvvrI8c51bDnK54PrDQ25Ow1R2xryEaelzS3+tRkGyFCs2AZO9hrRzvBv6cCgYEA1g0Wyvq8Dgbkm2DReCpmzSQRMfswF75uJ/kr+SIYV4OvZh9x4xrRnvOvVOw2eN5wmg7icdPInthRo7DV8N4AWwHWMTY5DQuZ3jFTgQjXHcZTloUl8hurDG4biR7WHLr3aWNSdohO0QsW1hD44gF+C3IaIzbFil4xqyTu1+veFskCgYEAtbZkwwU+aYVw0vGJV/r4BsKC6wbxePVW+R6qlmnoIXS4d5v5QIuBxFz2rqTHbM+/6Fu66fUHQyeUn+wm2Mm6UEEk4KfsKXt+oPcCQuiVFmUCNNRAU2g26cdMdwJdAeM1Ga1j1IKViz3d+V25tdzPUQTubCp1YMVxJGSeQ/nydHsCgYANqnKkfHYAdHPhJZbRfzXUJ4+oUIBW7XliZbc3eNc91rn/xi/WORjzbqvmFMqYiFDAfTQTIv05O7w9Gj1Zm1gchzAzwYwoCVbKumKJ54dc6pgO6zasU4m87CZOKFsBTMCVcNpYxAhEEOEczdTMeeIfNruesbjdCu6fOlkooZm3IQ==",
// 		ReturnUrl: "http://127.0.0.1:4200/display-data?title=%E9%99%88%E9%A1%BA%E8%A1%8C%E7%AD%89%E6%95%85%E6%84%8F%E4%BC%A4%E5%AE%B3%E6%A1%88",
// 		NotifyUrl: "http://127.0.0.1:4200/display-data?title=%E9%99%88%E9%A1%BA%E8%A1%8C%E7%AD%89%E6%95%85%E6%84%8F%E4%BC%A4%E5%AE%B3%E6%A1%88",
// 		Email:     "bcsbxs2618@sandbox.com",
// 	}

// 	form := Alipay.Form(alipay.Options{
// 		OrderId:  "123",
// 		Fee:      99.8,
// 		NickName: "翱翔大空",
// 		Subject:  "充值100",
// 	})

// 	if form == "" {
// 		fmt.Println("可以成功吗？")
// 	}

// 	var info string = "可以成功吗？"

// 	response := Response{info}

// 	json, _ := json.Marshal(response)

// 	w.Write(json)
// }
