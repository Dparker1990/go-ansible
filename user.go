package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

type User struct {
	username string
	conn     net.Conn
}

func (u *User) acquireUsername() {
	var username string
	fmt.Print("Please enter your username: ")
	if _, err := fmt.Scanln(&username); err != nil {
		log.Fatalf("Error trying to receive username. Failed with: %s", err.Error())
	}

	u.username = username
}

func (u *User) Connect() net.Conn {
	var err error

	u.acquireUsername()
	u.conn, err = net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("Client could not connect to server. Failed with: %s", err.Error())
	}
	u.SendMessage([]byte("Entered the room\n"))

	return u.conn
}

func (u User) SendMessage(msg []byte) {
	message := bufio.NewWriter(u.conn)

	message.WriteString(u.username + " > ")
	message.Write(msg)
	if err := message.Flush(); err != nil {
		log.Fatalf("Error when writting to room, failed with: %s", err.Error())
	}
}

func (u User) WriteUsername() {
	fmt.Printf("%s > ", u.username)
}

func (u User) WaitForInput() {
	QUIT := []byte("/quit\n")
	buf := bufio.NewReader(os.Stdin)

	for {
		u.WriteUsername()
		msg, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatal("Could not read user input")
		}

		if bytes.Equal(msg, QUIT) {
			u.SendMessage([]byte("left the room\n"))
			break
		}

		u.SendMessage(msg)
	}
}
