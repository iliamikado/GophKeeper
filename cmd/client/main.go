// Пакет с клиентом
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var client http.Client
var baseUrl = "http://localhost:8080/"
var jwt string
var BuildVersion string = "N/A"
var BuildDate string = "N/A"

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\n", BuildVersion, BuildDate)
	client = http.Client{}
	run()
}

func run() {
	r := bufio.NewReader(os.Stdin)
	out:
	for {
		fmt.Print(">>> ")
		str, _ := r.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		command := strings.Split(str, " ")
		switch command[0] {
			case "end":
				break out
			case "help":
				help()
			case "login":
				login(command)
			case "register":
				register(command)
			case "save_lp":
				save_l_p(command)
			case "get_lp":
				get_l_p(command)
			case "save_text":
				save_text(command)
			case "get_text":
				get_text(command)
			case "save_card":
				save_card(command)
			case "get_card":
				get_card(command)
			case "get_all":
				get_all()
			default:
				fmt.Println("Unexpected command. Type help to see all commands")
		}
	}
}

const helpDoc string = `
help - see all commands
end - close client
register [login] [password] - register new user
login [login] [password] - login in existing account
save_lp [login] [password] [metadata?] - save login and password
get_lp [key] - get login and password by key
save_text [text] [metadata?] - save text
get_text [key] - get text by key
save_card [number] [year_and_month] [cvv] [metadata?] - save login and password
get_card [key] - get card by key
get_all - get all saved data`

func help() {
	fmt.Println(helpDoc)
}

func sendReq(method, url string, body any) (*http.Response, []byte) {
	b, _ := json.Marshal(body)
	req, err := http.NewRequest(method, baseUrl + url, bytes.NewReader(b))
	req.AddCookie(&http.Cookie{
		Name: "JWT",
		Value: jwt,
	})
	if (err != nil) {
		fmt.Println(err.Error())
		return nil, nil
	}
	resp, err := client.Do(req)
	if (err != nil) {
		fmt.Println(err.Error())
		return nil, nil
	}
	defer resp.Body.Close()
	ans, _ := io.ReadAll(resp.Body)
	fmt.Println("Status code: " + resp.Status)
	if (resp.StatusCode != http.StatusOK) {
		return nil, nil
	}
	return resp, ans
}