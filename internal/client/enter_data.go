package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EnterData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

func (cl *Client) save_l_p(command []string) {
	if len(command) < 3 {
		fmt.Println("Not enough arguments")
		return
	}
	var enterData = EnterData{
		Login:    command[1],
		Password: command[2],
	}
	if len(command) > 3 {
		enterData.Metadata = command[3]
	}
	_, ans := cl.sendReq(http.MethodPost, "enter_data/", enterData)
	if ans != nil {
		fmt.Println("Data saved. The key is " + string(ans))
	}

}

type GetEnterDataReq struct {
	Key string `json:"key"`
}

func (cl *Client) get_l_p(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var getEnterDataReq = GetEnterDataReq{
		Key: command[1],
	}
	_, ans := cl.sendReq(http.MethodGet, "enter_data/", getEnterDataReq)
	if ans != nil {
		var enterData EnterData
		json.Unmarshal(ans, &enterData)
		fmt.Println("The data is:")
		fmt.Println("Login: " + enterData.Login)
		fmt.Println("Password: " + enterData.Password)
		fmt.Println("Metadata: " + enterData.Metadata)
	}
}
