package entities

type PokemonType struct {
	Id   int
	Name string

	IsSuperEffective []PokemonType
}

type Pokemon struct {
	Name    string
	TypeOne *PokemonType
	TypeTwo *PokemonType
}

type PokemonSocketData struct {
	Id int
}
