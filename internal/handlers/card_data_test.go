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

func TestSaveCardData(t *testing.T) {
	initApp()
	cardData := BankCardData{
		Number: "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV: "333",
		Metadata: "123",
	}
	req := httptest.NewRequest(http.MethodPost, "/save_card_data", dataToBody(cardData))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	SaveBankCardData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	key := rr.Body.String()
	assert.NotEmpty(t, key)
	data, err := passwordManager.GetBankCardData(defaultLogin, key)
	assert.NoError(t, err)
	assert.Equal(t, data.Number, cardData.Number)
	assert.Equal(t, data.YearAndMonth, cardData.YearAndMonth)
	assert.Equal(t, data.CVV, cardData.CVV)
	assert.Equal(t, data.Metadata, cardData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/save_card_data", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	SaveBankCardData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestGetCardData(t *testing.T) {
	initApp()
	cardData := models.BankCardData{
		Number: "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV: "333",
		Data: models.Data{
			Metadata: "123",
		},
	}
	key := passwordManager.SaveBankCardData(defaultLogin, cardData)

	req := httptest.NewRequest(http.MethodGet, "/get_card_data", dataToBody(GetBankCardDataReq{key}))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetBankCardData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData BankCardData
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, respData.Number, cardData.Number)
	assert.Equal(t, respData.YearAndMonth, cardData.YearAndMonth)
	assert.Equal(t, respData.CVV, cardData.CVV)
	assert.Equal(t, respData.Metadata, cardData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/get_card_data", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	GetBankCardData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}