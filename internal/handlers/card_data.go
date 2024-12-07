package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

func SavePaymentCard(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var paymentCard models.PaymentCard
	if err := json.NewDecoder(r.Body).Decode(&paymentCard); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SavePaymentCard(login, paymentCard)
	logger.Info("Save key " + key)
	w.Write([]byte(key))
}

type GetPaymentCardReq struct {
	Key string `json:"key"`
}

func GetPaymentCard(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var getPaymentCardReq GetPaymentCardReq
	if err := json.NewDecoder(r.Body).Decode(&getPaymentCardReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := getPaymentCardReq.Key
	paymentCard, err := passwordManager.GetPaymentCard(login, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(paymentCard)
	w.Write(resp)
}
