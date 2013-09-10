package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func listenForMessages(conn net.Conn) {
	msg := make([]byte, 1024)
	for {
		if _, err := conn.Read(msg); err != nil {
			log.Fatal("Could not read from user socket")
		}

		fmt.Println(string(msg))
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
