package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"junidex/entities"
	"log"
	"os"
	"path/filepath"
)

type PokemonJsonRepository struct {
	PokemonMap map[int]entities.Pokemon
	PokemonMapName map[string]entities.Pokemon
}

var repository PokemonJsonRepository

func GetJsonInstance() PokemonJsonRepository {
	if repository.PokemonMap == nil {
		fmt.Println("Setting up JSON repository")

		repository = InitJsonRepository()
		return repository
	}

	return repository
}

func GetDataFilePath(file string) string {
	return filepath.Join(os.Getenv("DATA_FOLDER"), file)
}

func InitJsonRepository() PokemonJsonRepository {
	pokemonMap := make(map[int]entities.Pokemon)
	pokemonMapName := make(map[string]entities.Pokemon)

	// Load json
	jsonFile, err := os.Open(GetDataFilePath("pokemon.json"))

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
		pokemonMapName[pokemon.Name] = pokemon
	}

	return PokemonJsonRepository{PokemonMap: pokemonMap, PokemonMapName: pokemonMapName}
}

func SaveTeam(teamMembers []string) {
	jsonString, _ := json.Marshal(teamMembers)

	log.Println("Saving Team locally...")
	ioutil.WriteFile(GetDataFilePath("team.json"), jsonString, os.ModePerm)
	log.Println("Team saved.")
}

func LoadTeam() []string {
	jsonFile, err := os.Open(GetDataFilePath("team.json"))

	if err != nil {
		return []string{}
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var teamMembers []string

	json.Unmarshal(byteValue, &teamMembers)

	return teamMembers
}
