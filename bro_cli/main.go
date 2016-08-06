package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BroList struct {
	bros []string
}

func main() {
	fmt.Println("Welcome to Bro CLI\nType help to get a list of valid commands")
	menu()
}

func menu() {
	for {
		var s string
		scnr := bufio.NewScanner(os.Stdin)
		done := scnr.Scan()
		if done {
			s = scnr.Text()
		}
		if strings.EqualFold(s, "help") {
			showCommands()
		}
		if strings.EqualFold(s, "brolist") {
			showBroList()
		}
	}
}

func showCommands() {
	fmt.Println("register --create a new bro account\n" +
		"login [username] [password] --already have a Bro account? Just Login\n" +
		"add [bro_name] --add a new bro\n" +
		"brolist --show a list of all your bros\n" +
		"bro [bro_name] --open a chat with one of your bros\n")
}

func showBroList() {
	resp, err := http.Get("") //get list
	return resp
}
