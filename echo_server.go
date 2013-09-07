package main

import (
	"log"
	"net"
)

func handleConnection(connection net.Conn, connections *[]net.Conn) {
	var conn net.Conn
	msg := make([]byte, 1024)

	for {
		if _, err := connection.Read(msg); err != nil {
			log.Printf("Connection closed: %s", connection.RemoteAddr().String())
			break
		}
		log.Printf("Message received: %s", string(msg))

		for i := range *connections {
			conn = (*connections)[i]
			if conn != connection {
				if _, err := conn.Write(msg); err != nil {
					log.Fatal("Writing failed")
				}
			}
		}
	}
}

func main() {
	connections := make([]net.Conn, 0, 100000)
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

		connections = append(connections, connection)

		go handleConnection(connection, &connections)
	}
}
