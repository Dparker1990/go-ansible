package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/wsxiaoys/terminal"
	"log"
	"net"
)

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

func runClient() {
	user := User{}
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
