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
	cardData1 := models.BankCardData{
		Number: "2200 0000 0000 0000",
		YearAndMonth: "08/08",
		CVV: "333",
		Data: models.Data{
			Metadata: "123",
		},
	}
	cardData2 := models.BankCardData{
		Number: "2200 2200 0000 0000",
		YearAndMonth: "08/08",
		CVV: "333",
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
	passwordManager.SaveBankCardData(defaultLogin, cardData1)
	passwordManager.SaveBankCardData(defaultLogin, cardData2)
	passwordManager.SaveTextData(defaultLogin, textData)

	req := httptest.NewRequest(http.MethodGet, "/get_all_data", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetAllData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData AllData
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, len(respData.CardData), 2)
	assert.Equal(t, len(respData.TextData), 1)
	assert.Equal(t, len(respData.EnterData), 0)
	assert.Equal(t, respData.CardData[0].Number, cardData1.Number)
	assert.Equal(t, respData.CardData[1].Number, cardData2.Number)
	assert.Equal(t, respData.TextData[0].Text, textData.Text)
}