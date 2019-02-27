package serve

import (
	"ChatServe/data"
	"ChatServe/utils"
	"encoding/json"
	"net/http"
)

var cookie http.Cookie

// ListThread 列出所有的主贴
func ListThread(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	threads, err := data.GetThread()
	if err != nil {
		utils.Logger.SetPrefix("ERROR ")
		utils.Logger.Println(err)
		return
	}
	data, err := json.Marshal(threads)
	if err != nil {
		utils.Logger.SetPrefix("ERROR ")
		utils.Logger.Println(err)
		return
	}
	w.Write(data)
}
