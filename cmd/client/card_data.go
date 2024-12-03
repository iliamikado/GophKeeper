package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BankCardData struct {
	Number string `json:"number"`
	YearAndMonth string `json:"year_and_month"`
	CVV string `json:"cvv"`
	Metadata string `json:"metadata"`
}

func save_card(command []string) {
	if len(command) < 4 {
		fmt.Println("Not enough arguments")
		return
	}
	var bankCardData = BankCardData{
		Number: command[1],
		YearAndMonth: command[2],
		CVV: command[3],
	}
	if len(command) > 4 {
		bankCardData.Metadata = command[4]
	}
	_, ans := sendReq(http.MethodPost, "save_card_data", bankCardData)
	if ans != nil {
		fmt.Println("Data saved. The key is " + string(ans))
	}

}

type GetBankCardDataReq struct {
	Key string `json:"key"`
}

func get_card(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var getBankCardDataReq = GetBankCardDataReq{
		Key: command[1],
	}
	_, ans := sendReq(http.MethodGet, "get_card_data", getBankCardDataReq)
	if ans != nil {
		var bankCardData BankCardData
		json.Unmarshal(ans, &bankCardData)
		fmt.Println("The data is:")
		fmt.Println("Number: " + bankCardData.Number)
		fmt.Println("Year and Month: " + bankCardData.YearAndMonth)
		fmt.Println("CVV: " + bankCardData.CVV)
		fmt.Println("Metadata: " + bankCardData.Metadata)
	}
}