package handlers

import (
	"PasswordManager/internal/logger"
	"PasswordManager/internal/models"
	"encoding/json"
	"net/http"
)

type PaymentCard struct {
	Number       string `json:"number"`
	YearAndMonth string `json:"year_and_month"`
	CVV          string `json:"cvv"`
	Metadata     string `json:"metadata"`
}

func SavePaymentCard(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	var paymentCard PaymentCard
	if err := json.NewDecoder(r.Body).Decode(&paymentCard); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := passwordManager.SavePaymentCard(login, models.PaymentCard{
		Number:       paymentCard.Number,
		YearAndMonth: paymentCard.YearAndMonth,
		CVV:          paymentCard.CVV,
		Data: models.Data{
			Metadata: paymentCard.Metadata,
		},
	})
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
	data, err := passwordManager.GetPaymentCard(login, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(PaymentCard{
		Number:       data.Number,
		YearAndMonth: data.YearAndMonth,
		CVV:          data.CVV,
		Metadata:     data.Metadata,
	})
	w.Write(resp)
}
