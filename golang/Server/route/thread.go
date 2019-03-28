package route

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"workplace/psql"
)

// ListThread 列出所有的主贴
func ListThread(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	threads, err := psql.GetThread()
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	data, err := json.Marshal(threads)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	w.Write(data)
}

// Search 讨论区搜索
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str := r.FormValue("condition")
	conditions := strings.Fields(str)
	thread, err := psql.RsByCondition(conditions)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	data, err := json.Marshal(thread)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	w.Write(data)
}

// Read 浏览帖子时，主贴的浏览数加一
func Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str := r.FormValue("topicID")
	topicID, err := strconv.Atoi(str)
	if err != nil {
		psql.Logger.SetPrefix("ERROR ")
		psql.Logger.Println(err)
		return
	}
	psql.AddVisitNum(topicID)
}
