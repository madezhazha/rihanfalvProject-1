package serve

import (
	"ChatServe/data"
	"ChatServe/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Search 讨论区搜索
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str := r.FormValue("condition")
	conditions := strings.Fields(str)
	fmt.Println(conditions)
	thread, err := data.RsByCondition(conditions)
	if err != nil {
		utils.Logger.SetPrefix("ERROR ")
		utils.Logger.Println(err)
		return
	}
	data, err := json.Marshal(thread)
	if err != nil {
		utils.Logger.SetPrefix("ERROR ")
		utils.Logger.Println(err)
		return
	}
	w.Write(data)
}
