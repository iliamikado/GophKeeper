// Пакет с клиентом
package client

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

var baseUrl = "http://localhost:8080/api/v1/"

type Client struct {
	cl http.Client
	jwt string
}

func StartClient() {
	appClient := Client{
		cl: http.Client{},
		jwt: "",
	}
	run(&appClient)
}

func run(client *Client) {
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
				client.login(command)
			case "register":
				client.register(command)
			case "save_lp":
				client.save_l_p(command)
			case "get_lp":
				client.get_l_p(command)
			case "save_text":
				client.save_text(command)
			case "get_text":
				client.get_text(command)
			case "save_card":
				client.save_card(command)
			case "get_card":
				client.get_card(command)
			case "get_all":
				client.get_all()
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

func (cl *Client) sendReq(method, url string, body any) (*http.Response, []byte) {
	b, _ := json.Marshal(body)
	req, err := http.NewRequest(method, baseUrl + url, bytes.NewReader(b))
	req.AddCookie(&http.Cookie{
		Name: "JWT",
		Value: cl.jwt,
	})
	if (err != nil) {
		fmt.Println(err.Error())
		return nil, nil
	}
	resp, err := cl.cl.Do(req)
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