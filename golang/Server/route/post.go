package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"workplace/psql"
)

// ListPost 列出某一主贴的所有回帖
func ListPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str1 := r.FormValue("topicID")
	str2 := r.FormValue("userID")
	topicID, err := strconv.Atoi(str1)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	userID, err := strconv.Atoi(str2)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 获取楼主信息
	thread, err := psql.ThreadByTopicID(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 获取所有回帖
	posts, err := psql.GetPost(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 判断当前登陆的用户是否收藏了此贴
	collection, err := psql.IsCollected(userID, topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	threadAndPost := make(map[string]interface{})
	threadAndPost["thread"] = thread
	threadAndPost["post"] = posts
	threadAndPost["collection"] = collection

	data, err := json.Marshal(threadAndPost)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	w.Write(data)
}

// CreatePost 新建一个新的回帖
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str1 := r.FormValue("userID")
	str2 := r.FormValue("topicID")
	text := r.FormValue("text")
	str4 := r.FormValue("floor")
	userID, err := strconv.Atoi(str1)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	topicID, err := strconv.Atoi(str2)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	floor, err := strconv.Atoi(str4)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 创建一个新的回帖
	err = psql.CreatePost(userID, topicID, text, floor)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 修改数据库主贴的回帖数
	err = psql.AddRepNum(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 更新主贴的最后回复时间
	err = psql.UpdFinReplyTime(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
}

// Collect 收藏主贴
func Collect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str1 := r.FormValue("userID")
	str2 := r.FormValue("topicID")
	userID, err := strconv.Atoi(str1)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	topicID, err := strconv.Atoi(str2)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 登陆用户收藏主贴
	collectionid, err := psql.Collect(userID, topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 有用户收藏主贴时，修改该主贴的收藏数
	err = psql.AddCollectNum(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	data, err := json.Marshal(collectionid)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	w.Write(data)
}

// Cancel 取消收藏
func Cancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str1 := r.FormValue("userID")
	str2 := r.FormValue("topicID")
	userID, err := strconv.Atoi(str1)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	topicID, err := strconv.Atoi(str2)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 登陆用户取消收藏主贴
	collectionid, err := psql.Cancel(userID, topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}

	// 每当有用户取消收藏主贴时，主贴的收藏数减一
	err = psql.CutCollectNum(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	data, err := json.Marshal(collectionid)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	w.Write(data)
}
