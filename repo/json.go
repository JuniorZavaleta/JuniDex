package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"junidex/entities"
	"log"
	"os"
)

type PokemonJsonRepository struct {
	PokemonMap map[int]entities.Pokemon
}

var repository PokemonJsonRepository

func GetJsonInstance() PokemonJsonRepository {
	if repository.PokemonMap == nil {
		fmt.Println("Setting up JSON repository")

		repository = PokemonJsonRepository{InitJsonRepository()}
		return repository
	}

	return repository
}

func InitJsonRepository() map[int]entities.Pokemon {
	pokemonMap := make(map[int]entities.Pokemon)

	// Load json
	jsonFile, err := os.Open("./data/pokemon.json")

	if err != nil {
		log.Panic("Error accesing json file.")
	} else {
		log.Println("Json for Pokemon loaded correctly.")
	}

	defer jsonFile.Close()

	// read the opened file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pokemonList []entities.Pokemon
	json.Unmarshal(byteValue, &pokemonList)
	for _, pokemon := range pokemonList {
		pokemonMap[pokemon.Id] = pokemon
	}

	return pokemonMap
}

func SaveTeam(teamMembers []string) {
	jsonString, _ := json.Marshal(teamMembers)

	log.Println("Saving Team locally...")
	ioutil.WriteFile("data/team.json", jsonString, os.ModePerm)
	log.Println("Team saved.")
}
