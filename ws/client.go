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

type Request struct {
	Body map[string][]interface{} `json:"body"`
}

type Response struct {
	Body map[string]interface{} `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Handle message from WebClient
		var incoming Request
		err := c.Conn.ReadJSON(&incoming)

		if err != nil {
			log.Println(err)
			return
		}

		response := &Response{Body: make(map[string]interface{})}

		if team, ok := incoming.Body["team"]; ok {
			HandleTeam(incoming, response, team)
		} else if starter, ok := incoming.Body["starterId"]; ok {
			HandleStarter(incoming, response, starter[0])
		}

		c.Pool.Broadcast <- *response
		fmt.Printf("Message Received: %+v\n", incoming)
	}
}

func HandleTeam(i Request, o *Response, team []interface{}) {
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
	o.Body["team"] = teamMembers
	repo.SaveTeam(teamMembers)
}

func HandleStarter(i Request, o *Response, starter interface{}) {
	var jsonRepo = repo.GetJsonInstance()

	pokemonId := int(starter.(float64))
	starterName := jsonRepo.PokemonMap[pokemonId].Name

	fmt.Printf("Starter: %v \n", starterName)
	o.Body["starterName"] = starterName
}
