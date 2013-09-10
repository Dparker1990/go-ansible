package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

func trimNewline(msg string) (trimmedMsg string) {
	trimmedMsg = msg[0 : len(msg)-1]
	return
}

func listenForMessages(conn net.Conn) {
	buf := bufio.NewReader(conn)
	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal("Could not read from user socket")
		}

		fmt.Println(trimNewline(msg))
	}
}

func runClient() {
	user := User{}
	conn := user.Connect()
	go listenForMessages(conn)
	user.WaitForInput()
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
	flag.BoolVar(&server, "server", false, "Whether to run as server or client (default)")

	flag.Parse()

	if server {
		runServer()
	} else {
		runClient()
	}
}
