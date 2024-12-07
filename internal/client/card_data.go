package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PaymentCard struct {
	Number       string `json:"number"`
	YearAndMonth string `json:"year_and_month"`
	CVV          string `json:"cvv"`
	Metadata     string `json:"metadata"`
}

func (cl *Client) save_card(command []string) {
	if len(command) < 4 {
		fmt.Println("Not enough arguments")
		return
	}
	var paymentCard = PaymentCard{
		Number:       command[1],
		YearAndMonth: command[2],
		CVV:          command[3],
	}
	if len(command) > 4 {
		paymentCard.Metadata = command[4]
	}
	_, ans := cl.sendReq(http.MethodPost, "card_data/", paymentCard)
	if ans != nil {
		fmt.Println("Data saved. The key is " + string(ans))
	}

}

type GetPaymentCardReq struct {
	Key string `json:"key"`
}

func (cl *Client) get_card(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var getPaymentCardReq = GetPaymentCardReq{
		Key: command[1],
	}
	_, ans := cl.sendReq(http.MethodGet, "card_data/", getPaymentCardReq)
	if ans != nil {
		var paymentCard PaymentCard
		json.Unmarshal(ans, &paymentCard)
		fmt.Println("The data is:")
		fmt.Println("Number: " + paymentCard.Number)
		fmt.Println("Year and Month: " + paymentCard.YearAndMonth)
		fmt.Println("CVV: " + paymentCard.CVV)
		fmt.Println("Metadata: " + paymentCard.Metadata)
	}
}
