package server

import (
	"bufio"
	"container/list"
	"github.com/Dparker1990/go-ansible/util"
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
				log.Println("Writing failed")
				log.Fatal(err)
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
		log.Printf("Message received: %s", util.TrimNewline(msg))

		writeToConnections(connections, connection, []byte(msg))
	}
}

func AcceptConnections(server net.Listener) {
	connChan := make(chan net.Conn, 10)
	defer close(connChan)

	go manageIncomingConnections(connChan)

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Println("Could not accept connection")
			log.Fatal(err)
			continue
		}
		log.Printf("Connection received from: %s", connection.RemoteAddr().String())

		connChan <- connection
	}
}
