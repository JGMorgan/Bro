package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
	"strings"
)

type BroList struct {
	bros []string
}

func main() {
	fmt.Println("Welcome to Bro CLI\nType help to get a list of valid commands")
	ws, err := websocket.Dial("ws://localhost:8000/socket", "", "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}
	for {
		message := []byte("hello, world!")
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent: %s\n", message)

		var msg = make([]byte, 512)
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s\n", msg)
	}
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
	http.Get("") //get list
}
