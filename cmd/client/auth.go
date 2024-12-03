package main

import (
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func login(command []string) {
	if len(command) < 3 {
		fmt.Println("Not enough arguments")
		return
	}
	var user = User{
		Login:    command[1],
		Password: command[2],
	}
	resp, _ := sendReq(http.MethodPost, "login", user)
	if resp != nil {
		jwt = strings.TrimPrefix(resp.Header.Get("Set-Cookie"), "JWT=")
	}
}

func register(command []string) {
	if len(command) < 3 {
		fmt.Println("Not enough arguments")
		return
	}
	var user = User{
		Login:    command[1],
		Password: command[2],
	}
	resp, _ := sendReq(http.MethodPost, "register", user)
	if resp != nil {
		jwt = strings.TrimPrefix(resp.Header.Get("Set-Cookie"), "JWT=")
	}
}
