package entities

type Evolution struct {
	Type      string `json:"type"`
	Evolution string `json:"evolution"`
}

type Pokemon struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	TypeOne    string      `json:"type1"`
	TypeTwo    string      `json:"type2"`
	Evolutions []Evolution `json:"evolutions"`
}
