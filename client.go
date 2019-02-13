package main

import (
	"github.com/gorilla/websocket"
)

//client represents a single chatting user.
type client struct {
	//socket web socket for this client
	socket *websocket.Conn
	//sent is a channel to send the messages
	send chan []byte
	//foom for the chat room to send messages
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
