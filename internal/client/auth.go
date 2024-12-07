package client

import (
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (cl *Client) login(command []string) {
	if len(command) < 3 {
		fmt.Println("Not enough arguments")
		return
	}
	var user = User{
		Login:    command[1],
		Password: command[2],
	}
	resp, _ := cl.sendReq(http.MethodPost, "login", user)
	if resp != nil {
		cl.jwt = strings.TrimPrefix(resp.Header.Get("Set-Cookie"), "JWT=")
	}
}

func (cl *Client) register(command []string) {
	if len(command) < 3 {
		fmt.Println("Not enough arguments")
		return
	}
	var user = User{
		Login:    command[1],
		Password: command[2],
	}
	resp, _ := cl.sendReq(http.MethodPost, "register", user)
	if resp != nil {
		cl.jwt = strings.TrimPrefix(resp.Header.Get("Set-Cookie"), "JWT=")
	}
}
