package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TextData struct {
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

func (cl *Client) save_text(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var textData = TextData{
		Text: command[1],
	}
	if len(command) > 2 {
		textData.Metadata = command[2]
	}
	_, ans := cl.sendReq(http.MethodPost, "text_data/save", textData)
	if ans != nil {
		fmt.Println("Data saved. The key is " + string(ans))
	}

}

type GetTextDataReq struct {
	Key string `json:"key"`
}

func (cl *Client) get_text(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var getTextDataReq = GetTextDataReq{
		Key: command[1],
	}
	_, ans := cl.sendReq(http.MethodGet, "text_data/get", getTextDataReq)
	if ans != nil {
		var textData TextData
		json.Unmarshal(ans, &textData)
		fmt.Println("The data is:")
		fmt.Println("Text: " + textData.Text)
		fmt.Println("Metadata: " + textData.Metadata)
	}
}