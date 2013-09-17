package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/wsxiaoys/terminal"
	"log"
	"net"
)

func promtForUserName() (username string) {
	fmt.Print("Please enter your username: ")
	if _, err := fmt.Scanln(&username); err != nil {
		log.Fatalf("Error trying to receive username. Failed with: %s", err.Error())
	}

	return
}

func writeMessage(msg string, user User) {
	terminal.Stdout.ClearLine()
	terminal.Stdout.Left(50)
	fmt.Println(trimNewline(msg))
	user.WriteUsername()
}

func listenForMessages(conn net.Conn, user User) {
	buf := bufio.NewReader(conn)
	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Println("Could not read from user socket")
			log.Fatal(err)
		}

		writeMessage(msg, user)
	}
}

func acquireUsername(user *User) {
	username := promtForUserName()
	if err := user.SetUsername(username); err != nil {
		fmt.Println(err)
		acquireUsername(user)
	}
}

func runClient() {
	user := User{}
	acquireUsername(&user)
	conn := user.Connect()
	go listenForMessages(conn, user)
	user.WaitForInput()
	conn.Close()
}

func runServer() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Could not listen")
		log.Fatal(err)
	}
	defer server.Close()
	acceptConnections(server)
}

func main() {
	var server bool
	flag.BoolVar(&server, "s", false, "Whether to run as server or client (default)")

	flag.Parse()

	if server {
		runServer()
	} else {
		runClient()
	}
}
