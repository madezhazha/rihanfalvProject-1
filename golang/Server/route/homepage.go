package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../psql"
)

func GetHomePageArtical(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		return
	}
	var postbody map[string]int64
	json.Unmarshal(body, &postbody)
	index := postbody["index"] //get artical index from which
	var date = psql.GetHPADate(index)
	date_json, _ := json.Marshal(date)
	w.Write(date_json)
}

func GetHomePageHotnews(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	date := psql.GetHomePageHotnewDate()
	date_json, _ := json.Marshal(date)
	w.Write(date_json)
}
