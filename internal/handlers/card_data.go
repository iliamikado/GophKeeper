package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

type BankCardData struct {
	Number string `json:"number"`
	YearAndMonth string `json:"year_and_month"`
	CVV string `json:"cvv"`
	Metadata string `json:"metadata"`
}

func SaveBankCardData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var bankCardData BankCardData
	if err := json.NewDecoder(r.Body).Decode(&bankCardData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SaveBankCardData(login, models.BankCardData{
		Number: bankCardData.Number,
		YearAndMonth: bankCardData.YearAndMonth,
		CVV: bankCardData.CVV, 
		Data: models.Data{
			Metadata: bankCardData.Metadata,
		},
	})
	logger.Info("Save key " + key)
	w.Write([]byte(key))
}

type GetBankCardDataReq struct {
	Key string `json:"key"`
}

func GetBankCardData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var getBankCardDataReq GetBankCardDataReq
	if err := json.NewDecoder(r.Body).Decode(&getBankCardDataReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := getBankCardDataReq.Key
	data, err := passwordManager.GetBankCardData(login, key)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(BankCardData{
		Number: data.Number,
		YearAndMonth: data.YearAndMonth,
		CVV: data.CVV,
		Metadata: data.Metadata,
	})
	w.Write(resp)
}