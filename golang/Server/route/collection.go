package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../psql"
)

func UserCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm() //解析参数，默认是不会解析的

	RUserID, _ := ioutil.ReadAll(r.Body)
	var IUserID map[string]interface{}
	json.Unmarshal([]byte(RUserID), &IUserID)
	fmt.Printf("%v", IUserID)

	if IUserID != nil {

		InterUserID := IUserID["userid"]

		SUserID := InterUserID.(float64)

		fmt.Println(SUserID)

		var UserID int
		UserID = int(SUserID)

		fmt.Println(UserID)

		var mCollectionsMsg []psql.CollectionsMsg
		mCollectionsMsg = psql.SqlSelect(UserID)

		JmCollectionsMsg, err2 := json.Marshal(mCollectionsMsg)

		if err2 != nil {
			panic(err2)
		}

		w.Write([]byte(JmCollectionsMsg))
	}

}
