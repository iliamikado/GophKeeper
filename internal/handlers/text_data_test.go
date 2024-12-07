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

func TestSaveTextData(t *testing.T) {
	initApp()
	textData := TextData{
		Text:     "Hello",
		Metadata: "123",
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/text_data/", dataToBody(textData))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	SaveTextData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	key := rr.Body.String()
	assert.NotEmpty(t, key)
	data, err := passwordManager.GetTextData(defaultLogin, key)
	assert.NoError(t, err)
	assert.Equal(t, data.Text, textData.Text)
	assert.Equal(t, data.Metadata, textData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/text_data/", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	SaveTextData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestGetTextData(t *testing.T) {
	initApp()
	textData := models.TextData{
		Text: "Hello",
		Data: models.Data{
			Metadata: "123",
		},
	}
	key := passwordManager.SaveTextData(defaultLogin, textData)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/text_data/", dataToBody(GetEnterDataReq{key}))
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr := httptest.NewRecorder()

	GetTextData(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	respBody, _ := io.ReadAll(rr.Body)
	var respData TextData
	json.Unmarshal(respBody, &respData)
	assert.NotEmpty(t, respData)
	assert.Equal(t, respData.Text, textData.Text)
	assert.Equal(t, respData.Metadata, textData.Metadata)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/text_data/", nil)
	req = req.WithContext(context.WithValue(req.Context(), loginKey{}, defaultLogin))
	rr = httptest.NewRecorder()
	GetTextData(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}
