package main

import (
	"PasswordManager/internal/client"
	"fmt"
)

var BuildVersion string = "N/A"
var BuildDate string = "N/A"

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\n", BuildVersion, BuildDate)
	client.StartClient()
}