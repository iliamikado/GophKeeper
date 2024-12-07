package handlers

import (
	"PasswordManager/internal/models"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSavePaymentCard(t *testing.T) {
	initApp()
	cardData := PaymentCard{
		Number:       "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV:          "333",
		Metadata:     "123",
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/payment_card/", dataToBody(cardData))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	SavePaymentCard(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	key := rr.Body.String()
	assert.NotEmpty(t, key)
	data, err := passwordManager.GetPaymentCard(defaultLogin, key)
	assert.NoError(t, err)
	assert.Equal(t, data.Number, cardData.Number)
	assert.Equal(t, data.YearAndMonth, cardData.YearAndMonth)
	assert.Equal(t, data.CVV, cardData.CVV)
	assert.Equal(t, data.Metadata, cardData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/payment_card/", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	SavePaymentCard(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestGetPaymentCard(t *testing.T) {
	initApp()
	cardData := models.PaymentCard{
		Number:       "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV:          "333",
		Data: models.Data{
			Metadata: "123",
		},
	}
	key := passwordManager.SavePaymentCard(defaultLogin, cardData)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/payment_card/", dataToBody(GetPaymentCardReq{key}))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetPaymentCard(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData PaymentCard
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, respData.Number, cardData.Number)
	assert.Equal(t, respData.YearAndMonth, cardData.YearAndMonth)
	assert.Equal(t, respData.CVV, cardData.CVV)
	assert.Equal(t, respData.Metadata, cardData.Metadata)

	req = httptest.NewRequest(http.MethodGet, "/api/v1/payment_card/", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	GetPaymentCard(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}
