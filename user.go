package main

import (
	"bufio"
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
		log.Fatal("Error trying to receive username")
	}

	u.username = username
}

func (u *User) Connect() net.Conn {
	var err error

	u.acquireUsername()
	u.conn, err = net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("Client could not connect to server.")
	}
	u.SendMessage([]byte(u.username + " has entered the room\n"))

	return u.conn
}

func (u User) SendMessage(msg []byte) {
	if _, err := u.conn.Write(msg); err != nil {
		log.Fatal("Error when writting to room")
	}
}

func (u User) WaitForInput() {
	buf := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s > ", u.username)
		msg, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal("Could not read user input")
		}

		if msg == "/quit\n" {
			break
		}

		u.SendMessage([]byte(msg))
	}
}
