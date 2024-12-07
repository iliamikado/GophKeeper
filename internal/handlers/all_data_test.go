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

func TestGetAllData(t *testing.T) {
	initApp()
	cardData1 := models.PaymentCard{
		Number:       "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV:          "333",
		Data: models.Data{
			Metadata: "123",
		},
	}
	cardData2 := models.PaymentCard{
		Number:       "2200 2200 0000 0000",
		YearAndMonth: "08/08",
		CVV:          "333",
		Data: models.Data{
			Metadata: "123",
		},
	}
	textData := models.TextData{
		Text: "abc",
		Data: models.Data{
			Metadata: "123",
		},
	}
	passwordManager.SavePaymentCard(defaultLogin, cardData1)
	passwordManager.SavePaymentCard(defaultLogin, cardData2)
	passwordManager.SaveTextData(defaultLogin, textData)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/get_all", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetAllData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData models.AllData
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, len(respData.PaymentCard), 2)
	assert.Equal(t, len(respData.TextData), 1)
	assert.Equal(t, len(respData.EnterData), 0)
	assert.Equal(t, respData.PaymentCard[0].Number, cardData1.Number)
	assert.Equal(t, respData.PaymentCard[1].Number, cardData2.Number)
	assert.Equal(t, respData.TextData[0].Text, textData.Text)
}
