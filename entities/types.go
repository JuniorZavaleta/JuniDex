package entities

const WATER_ID = 1
const GRASS_ID = 2
const FIRE_ID = 3

func WaterType() PokemonType {
	return PokemonType{Id: WATER_ID, Name: "Water"}
}

func GrassType() PokemonType {
	return PokemonType{Id: GRASS_ID, Name: "Grass"}
}

func FireType() PokemonType {
	return PokemonType{Id: FIRE_ID, Name: "Fire"}
}
