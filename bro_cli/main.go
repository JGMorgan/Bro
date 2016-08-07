package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/websocket"
	"log"
	//"math/rand"
	"net/http"
	"neural_network"
	"os"
	//	"strconv"
	"math"
	"strings"
	//	"time"
)

type BroList struct {
	bros []string
}

func main() {
	fmt.Println("Welcome to Bro CLI\nSend a Bro to Bro Bot")
	//s := "CREATE TABLE bros (recvLength, deltaTime, respTime)"
	db, _ := sql.Open("sqlite3", "./test.db")
	//db.Exec(s)
	//stmt, _ := db.Prepare("INSERT INTO bros(recvLength,deltaTime ,respLength) values(?,?,?)")
	/*	s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		for i := 0; i < 1000; i++ {
			var buffer bytes.Buffer

			x := r1.Intn(20) + 1
			buffer.WriteString("INSERT INTO bros(recvLength, deltaTime, respTime) values(")
			buffer.WriteString(strconv.Itoa(x))
			buffer.WriteString(",")
			buffer.WriteString(strconv.Itoa(r1.Intn(5000) + 500))
			buffer.WriteString(",")
			buffer.WriteString(strconv.Itoa(x + r1.Intn(5) - 2))
			db.Exec(buffer.String())
			fmt.Println(i)
		}TEST DATA GENERATION*/

	net := neural_network.CreateNeuralNet(2, 1, 3)
	rows, _ := db.Query("SELECT * FROM bros")
	X := [][]float64{}
	Y := []float64{}
	for rows.Next() {
		var recvLength int
		var deltaTime int
		var respLength int
		rows.Scan(&recvLength, &deltaTime, &respLength)
		xRow := []float64{1 / float64(recvLength), 1 / float64(deltaTime)}
		X = append(X, xRow)
		Y = append(Y, 1/float64(respLength))
	}
	net.Train(X, Y, .1)
	recv := 1 / 1.0
	delta := 1 / 1.0
	resp := 1 / 1.0
	In := []float64{recv, delta}
	for {
		var s string
		scnr := bufio.NewScanner(os.Stdin)
		done := scnr.Scan()
		if done {
			s = scnr.Text()
		}
		recv = 1 / float64(len(s))
		In[0] = recv
		In[1] = recv

		resp = math.Ceil(1 / float64(net.Predict(In)))
		var buffer bytes.Buffer
		buffer.WriteString("br")
		for i := 0; i < int(resp); i++ {
			buffer.WriteString("o")
		}
		fmt.Println(buffer.String())

	}
}

func websocketHandler() {
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
