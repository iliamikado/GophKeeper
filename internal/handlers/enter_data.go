package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

type EnterData struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

func SaveEnterData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var enterData EnterData
	if err := json.NewDecoder(r.Body).Decode(&enterData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SaveEnterData(login, models.EnterData{
		Login: enterData.Login,
		Password: enterData.Password,
		Data: models.Data{
			Metadata: enterData.Metadata,
		},
	})
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
	data, err := passwordManager.GetEnterData(login, key)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(EnterData{
		Login: data.Login,
		Password: data.Password,
		Metadata: data.Metadata,
	})
	w.Write(resp)
}