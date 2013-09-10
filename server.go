package main

import (
	"bufio"
	"container/list"
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
		if buf[i] == 0x00 {
			break
		} else {
			buf[i] = 0x00
		}
	}
}

func manageIncomingConnections(connChan chan net.Conn) {
	connections := list.New()
	for connection := range connChan {
		connections.PushBack(connection)
		go handleConnection(connection, connections)
	}
}

func writeToConnections(connections *list.List, connection net.Conn, msg []byte) {
	for c := (*connections).Front(); c != nil; c = c.Next() {
		conn := c.Value.(net.Conn)
		if conn != connection {
			if _, err := conn.Write(msg); err != nil {
				log.Fatal("Writing failed")
			}
		}
	}
}

func handleConnection(connection net.Conn, connections *list.List) {
	buf := bufio.NewReader(connection)

	for {
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Printf("Connection closed: %s", connection.RemoteAddr().String())
			connection.Close()
			removeConnection(connections, connection)
			break
		}
		log.Printf("Message received: %s", trimNewline(msg))

		writeToConnections(connections, connection, []byte(msg))
	}
}

func acceptConnections(server net.Listener) {
	connChan := make(chan net.Conn, 10)
	go manageIncomingConnections(connChan)

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal("Could not accept connection")
			continue
		}
		log.Printf("Connection received from: %s", connection.RemoteAddr().String())

		connChan <- connection
	}
}
