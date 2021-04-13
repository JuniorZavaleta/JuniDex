package handlers

import (
	"html/template"
	"junidex/entities"
	"junidex/helpers"
	"junidex/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	// can add services
}

type PokemonTeam struct {
}

type AdminTeamPage struct {
	Team []entities.Pokemon
	PokemonList []entities.Pokemon
}

func NewAdminHandler(r *mux.Router) {
	handler := &AdminHandler{}

	r.HandleFunc("/admin/starter", handler.StarterView).Methods("GET")
	r.HandleFunc("/admin/team", handler.TeamView).Methods("GET")
}

func (h *AdminHandler) StarterView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("admin", "starter.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}

func (h *AdminHandler) TeamView(w http.ResponseWriter, r *http.Request) {
	teamMembers := repo.LoadTeam()
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("admin", "team.html"))

	if err != nil {
		log.Panic(err)
	}

	var team []entities.Pokemon
	var pokemonList []entities.Pokemon

	for _, pokemon := range repo.GetJsonInstance().PokemonMap {
		pokemonList = append(pokemonList, pokemon)
	}

	for _, pokemon := range teamMembers {
		team = append(team, repo.GetJsonInstance().PokemonMapName[pokemon])
	}

	p := AdminTeamPage{Team: team, PokemonList: pokemonList}

	t.Execute(w, p)
}
