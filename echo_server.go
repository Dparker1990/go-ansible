package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/wsxiaoys/terminal"
	"log"
	"net"
)

func trimNewline(msg string) (trimmedMsg string) {
	trimmedMsg = msg[0 : len(msg)-1]
	return
}

func listenForMessages(conn net.Conn, user User) {
	buf := bufio.NewReader(conn)
	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal("Could not read from user socket")
		}

		terminal.Stdout.ClearLine()
		terminal.Stdout.Left(50)
		fmt.Println(trimNewline(msg))
		fmt.Printf("%s > ", user.username)
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
		log.Fatal("Could not listen")
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
