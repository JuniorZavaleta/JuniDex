package handlers

import (
	"html/template"
	"junidex/helpers"
	"junidex/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ClientHandler struct {
	// can add services
}

type ClientTeamPage struct {
	Team []string
}

func NewClientHandler(r *mux.Router) {
	handler := &ClientHandler{}

	r.HandleFunc("/app/starter", handler.starterView).Methods("GET")
	r.HandleFunc("/app/team", handler.teamView).Methods("GET")
}

func (h *ClientHandler) starterView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("client", "starter.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}

func (h *ClientHandler) teamView(w http.ResponseWriter, r *http.Request) {
	// TODO: Should load more data about pokemon team - for now just names
	// team := &PokemonTeam{}

	team := repo.LoadTeam()
	log.Printf("Team loaded: %v", team)
	if len(team) == 0 {
		// If team not found, set 6 blank strings
		team = []string{"", "", "", "", "", ""}
	}

	p := ClientTeamPage{Team: team}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("client", "team.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, p)
}
