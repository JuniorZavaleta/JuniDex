package entities

type Pokemon struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	TypeOne string `json:"type1"`
	TypeTwo string `json:"type2"`
}

type PokemonSocketData struct {
	Id int
}
