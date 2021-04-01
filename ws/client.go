package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int         `json:"type"`
	Body interface{} `json:"body"`
}

type PokemonStarter struct {
	StarterId int `json:"starterId"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Handle message from WebClient
		var starter PokemonStarter
		err := c.Conn.ReadJSON(&starter)

		if err != nil {
			log.Println(err)
			return
		}

		// handle incoming message to get an appropiate outcoming message

		message := Message{Body: starter}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
