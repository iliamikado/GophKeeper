package handlers

import (
	"encoding/json"
	"net/http"
)

type AllData struct {
	EnterData []EnterData
	TextData []TextData
	CardData []BankCardData
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value(loginKey{}).(string)
	data := passwordManager.GetAllData(login)
	var allData = AllData{
		EnterData: make([]EnterData, len(data.EnterData)),
		TextData: make([]TextData, len(data.TextData)),
		CardData: make([]BankCardData, len(data.BankCardData)),
	}
	for i, x := range data.EnterData {
		allData.EnterData[i] = EnterData{
			Login: x.Login,
			Password: x.Password,
			Metadata: x.Metadata,
		}
	}
	for i, x := range data.TextData {
		allData.TextData[i] = TextData{
			Text: x.Text,
			Metadata: x.Metadata,
		}
	}
	for i, x := range data.BankCardData {
		allData.CardData[i] = BankCardData{
			Number: x.Number,
			YearAndMonth: x.YearAndMonth,
			CVV: x.CVV,
			Metadata: x.Metadata,
		}
	}
	resp, _ := json.Marshal(allData)
	w.Write(resp)
}