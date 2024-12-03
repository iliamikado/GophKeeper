package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TextData struct {
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

func save_text(command []string) {
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
	_, ans := sendReq(http.MethodPost, "save_text_data", textData)
	if ans != nil {
		fmt.Println("Data saved. The key is " + string(ans))
	}

}

type GetTextDataReq struct {
	Key string `json:"key"`
}

func get_text(command []string) {
	if len(command) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	var getTextDataReq = GetTextDataReq{
		Key: command[1],
	}
	_, ans := sendReq(http.MethodGet, "get_text_data", getTextDataReq)
	if ans != nil {
		var textData TextData
		json.Unmarshal(ans, &textData)
		fmt.Println("The data is:")
		fmt.Println("Text: " + textData.Text)
		fmt.Println("Metadata: " + textData.Metadata)
	}
}