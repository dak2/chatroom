package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行なっている1人のユーザーを表します
type client struct {
	// socketはこのクライアントのためのwebsocketです
	socket *websocket.Conn

	// sendはメッセージが送られるチャネルです
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		}else{
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send
}