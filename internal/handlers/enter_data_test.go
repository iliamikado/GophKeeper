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

func TestSaveEnterData(t *testing.T) {
	initApp()
	enterData := EnterData{
		Login: "aaa",
		Password: "bbb",
		Metadata: "123",
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/enter_data/save", dataToBody(enterData))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	SaveEnterData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	key := rr.Body.String()
	assert.NotEmpty(t, key)
	data, err := passwordManager.GetEnterData(defaultLogin, key)
	assert.NoError(t, err)
	assert.Equal(t, data.Login, enterData.Login)
	assert.Equal(t, data.Password, enterData.Password)
	assert.Equal(t, data.Metadata, enterData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/enter_data/save", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	SaveEnterData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestGetEnterData(t *testing.T) {
	initApp()
	enterData := models.EnterData{
		Login: "aaa",
		Password: "bbb",
		Data: models.Data{
			Metadata: "123",
		},
	}
	key := passwordManager.SaveEnterData(defaultLogin, enterData)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/enter_data/get", dataToBody(GetEnterDataReq{key}))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetEnterData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData EnterData
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, respData.Login, enterData.Login)
	assert.Equal(t, respData.Password, enterData.Password)
	assert.Equal(t, respData.Metadata, enterData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/enter_data/get", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	GetEnterData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}