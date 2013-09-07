package main

import (
	"container/list"
	"log"
	"net"
)

func handleConnection(connection net.Conn, connections *list.List) {
	var conn net.Conn
	msg := make([]byte, 1024)

	for {
		if _, err := connection.Read(msg); err != nil {
			log.Printf("Connection closed: %s", connection.RemoteAddr().String())
			break
		}
		log.Printf("Message received: %s", string(msg))

		for c := (*connections).Front(); c != nil; c = c.Next() {
			conn = c.Value.(net.Conn)
			if conn != connection {
				if _, err := conn.Write(msg); err != nil {
					log.Fatal("Writing failed")
				}
			}
		}
	}
}

func main() {
	connections := list.New()
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not listen")
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal("Could not accept connection")
			continue
		}
		log.Printf("Connection received from: %s", connection.RemoteAddr().String())

		connections.PushBack(connection)

		go handleConnection(connection, connections)
	}
}
