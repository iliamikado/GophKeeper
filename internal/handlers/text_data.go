package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

type TextData struct {
	Text string `json:"text"`
	Metadata string `json:"metadata"`
}

func SaveTextData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var textData TextData
	if err := json.NewDecoder(r.Body).Decode(&textData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SaveTextData(login, models.TextData{
		Text: textData.Text, 
		Data: models.Data{
			Metadata: textData.Metadata,
		},
	})
	logger.Info("Save key " + key)
	w.Write([]byte(key))
}

type GetTextDataReq struct {
	Key string `json:"key"`
}

func GetTextData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var getTextDataReq GetTextDataReq
	if err := json.NewDecoder(r.Body).Decode(&getTextDataReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := getTextDataReq.Key
	data, err := passwordManager.GetTextData(login, key)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(TextData{
		Text: data.Text,
		Metadata: data.Metadata,
	})
	w.Write(resp)
}