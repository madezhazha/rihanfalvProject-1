package route

import (
	//"ChatServe/psql"
	"encoding/json"
	"net/http"
	"strconv"

	"../psql"
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
	thread, err := psql.ThreadByTopicID(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	posts, err := psql.GetPost(topicID)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
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
	err = psql.CreatePost(userID, topicID, text, floor)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	err = psql.AddRepNum(topicID)
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
	collectionid, err := psql.Collect(userID, topicID)
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
	collectionid, err := psql.Cancel(userID, topicID)
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
