package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

func SaveEnterData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var enterData models.EnterData
	if err := json.NewDecoder(r.Body).Decode(&enterData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SaveEnterData(login, enterData)
	logger.Info("Save key " + key)
	w.Write([]byte(key))
}

type GetEnterDataReq struct {
	Key string `json:"key"`
}

func GetEnterData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var getEnterDataReq GetEnterDataReq
	if err := json.NewDecoder(r.Body).Decode(&getEnterDataReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := getEnterDataReq.Key
	enterData, err := passwordManager.GetEnterData(login, key)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(enterData)
	w.Write(resp)
}