package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

const maxUsernameLength int = 50

type User struct {
	username string
	conn     net.Conn
	msgbuf   *bufio.Writer
}

func (u *User) SetUsername(username string) error {
	if len(username) > maxUsernameLength {
		return fmt.Errorf("Username must be less than 50 characters")
	}
	u.username = username
	return nil
}

func (u *User) Connect() net.Conn {
	var err error
	u.conn, err = net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("Client could not connect to server. Failed with: %s", err.Error())
	}
	u.msgbuf = bufio.NewWriter(u.conn)
	u.SendMessage([]byte("Entered the room\n"))

	return u.conn
}

func (u *User) SendMessage(msg []byte) {
	u.msgbuf.WriteString(u.username + " > ")
	u.msgbuf.Write(msg)
	if err := u.msgbuf.Flush(); err != nil {
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
