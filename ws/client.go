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

type StarterIncoming struct {
	StarterId int `json:"starterId"`
}

type StarterOutcoming struct {
	StarterName string `json:"starterName"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Handle message from WebClient
		var starter StarterIncoming
		err := c.Conn.ReadJSON(&starter)

		if err != nil {
			log.Println(err)
			return
		}

		// handle incoming message to get an appropiate outcoming message
		response := &StarterOutcoming{}

		switch starter.StarterId {
			case 1:
				response.StarterName = "Bulbasaur"
			case 4:
				response.StarterName = "Squirtle"
			case 7:
				response.StarterName = "Charmander"
		}

		message := Message{Body: response}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
