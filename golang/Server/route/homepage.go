package route

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../psql"
)

func GetHomePageArtical(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	indexstr := vars["index"][0]
	index, _ := strconv.Atoi(indexstr)
	var date = psql.GetHPADate(index)
	date_json, _ := json.Marshal(date)
	w.Write(date_json)
}

func GetHomePageHotnews(w http.ResponseWriter, r *http.Request) {
	SetHeader(w)
	if r.Method != "GET" {
		return
	}
	date := psql.GetHomePageHotnewDate()
	date_json, _ := json.Marshal(date)
	w.Write(date_json)
}
