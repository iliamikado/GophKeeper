package handlers

import (
	"encoding/json"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	allData := passwordManager.GetAllData(login)
	resp, _ := json.Marshal(allData)
	w.Write(resp)
}
