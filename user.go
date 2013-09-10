package main

import (
	"fmt"
	"log"
	"net"
)

type User struct {
	username string
	conn     net.Conn
}

func (u *User) acquireUsername() {
	var username string
	fmt.Print("Please enter your username: ")
	if _, err := fmt.Scanln(&username); err != nil {
		log.Fatal("Error trying to receive username")
	}

	u.username = username
}

func (u *User) Connect() (conn net.Conn) {
	u.acquireUsername()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("Client could not connect to server.")
	}
	u.conn = conn
	u.SendMessage([]byte(u.username + " has entered the room"))

	return
}

func (u User) SendMessage(msg []byte) {
	if _, err := u.conn.Write(msg); err != nil {
		log.Fatal("Error when writting to room")
	}
}

func (u User) WaitForInput() {
	var msg string

	for {
		fmt.Printf("%s > ", u.username)
		fmt.Scanln(&msg)
		if msg == "/quit" {
			break
		}
		u.SendMessage([]byte(msg))
	}
}
