package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AllData struct {
	EnterData   []EnterData
	TextData    []TextData
	PaymentCard []PaymentCard
}

func get_all() {
	_, ans := sendReq(http.MethodGet, "get_all_data", nil)
	if ans != nil {
		var allData AllData
		json.Unmarshal(ans, &allData)
		fmt.Println("Total saved data - ", len(allData.PaymentCard)+len(allData.EnterData)+len(allData.TextData))
		fmt.Println("Enter data:")
		for _, data := range allData.EnterData {
			fmt.Println("Login: " + data.Login)
			fmt.Println("Password: " + data.Password)
			fmt.Println()
		}
		fmt.Println("Text data:")
		for _, data := range allData.TextData {
			fmt.Println("Text: " + data.Text)
			fmt.Println()
		}
		fmt.Println("Card data:")
		for _, data := range allData.PaymentCard {
			fmt.Println("Number: " + data.Number)
			fmt.Println("Year and month: " + data.YearAndMonth)
			fmt.Println("CVV: " + data.CVV)
			fmt.Println()
		}
	}
}
