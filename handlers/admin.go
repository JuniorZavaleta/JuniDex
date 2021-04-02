package handlers

import (
	"html/template"
	"junidex/helpers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	// can add services
}

type PokemonTeam struct {
}

func NewAdminHandler(r *mux.Router) {
	handler := &AdminHandler{}

	r.HandleFunc("/admin/starter", handler.adminView).Methods("GET")
	r.HandleFunc("/admin/team", handler.teamView).Methods("GET")
}

func (h *AdminHandler) adminView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("admin", "starter.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}

func (h *AdminHandler) teamView(w http.ResponseWriter, r *http.Request) {
	// TODO: Should load team in database
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("admin", "team.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}
