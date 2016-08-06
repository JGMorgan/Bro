package main

import (
	//	"bufio"
	"fmt"
	//	"os"
)

func main() {
	fmt.Println("Welcome to Bro CLI\nType help to get a list of valid commands")
	menu()
}

func menu() {
	for {
		var s string
		fmt.Scanln(&s)
		fmt.Println(s)
	}
}
