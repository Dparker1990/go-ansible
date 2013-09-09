package main

import (
	"container/list"
	"flag"
	"log"
	"net"
)

func removeConnection(connections *list.List, connection net.Conn) {
	for c := (*connections).Front(); c != nil; c = c.Next() {
		if c.Value.(net.Conn) == connection {
			(*connections).Remove(c)
		}
	}
}

func zeroBuffer(buf []byte) {
	for i := range buf {
		buf[i] = byte(0)
	}
}

func handleConnection(connection net.Conn, connections *list.List) {
	var conn net.Conn
	msg := make([]byte, 1024)

	for {
		zeroBuffer(msg)
		if _, err := connection.Read(msg); err != nil {
			log.Printf("Connection closed: %s", connection.RemoteAddr().String())
			connection.Close()
			removeConnection(connections, connection)
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

func runServer() {
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

func runClient() {
	user := User{}
	user.Connect()
	user.WaitForInput()
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
