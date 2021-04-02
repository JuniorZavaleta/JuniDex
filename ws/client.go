package ws

import (
	"fmt"
	"junidex/repo"
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

type RequestIncoming struct {
	Body map[string][]interface{} `json:"body"`
}

type RequestOutcoming struct {
	Body map[string]interface{} `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Handle message from WebClient
		var incoming RequestIncoming
		err := c.Conn.ReadJSON(&incoming)

		if err != nil {
			log.Println(err)
			return
		}

		response := &RequestOutcoming{Body: make(map[string]interface{})}

		if team, ok := incoming.Body["team"]; ok {
			var jsonRepo = repo.GetJsonInstance()

			teamMembers := []string{}
			for _, v := range(team) {
				// Idk why is coming as float64 number ._.
				pokemonId := int(v.(float64))

				if pokemon, ok := jsonRepo.PokemonMap[pokemonId]; ok {
					teamMembers = append(teamMembers, pokemon.Name)
				} else {
					teamMembers = append(teamMembers, "")
				}
			}

			fmt.Printf("Team: %v \n", teamMembers)
			response.Body["team"] = teamMembers
		}

		c.Pool.Broadcast <- *response
		fmt.Printf("Message Received: %+v\n", incoming)
	}
}
