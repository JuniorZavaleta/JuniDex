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

	r.HandleFunc("/admin", handler.adminView).Methods("GET")
}

func (h *AdminHandler) adminView(w http.ResponseWriter, r *http.Request) {
	team := &PokemonTeam{}
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("admin", "pokemon.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, team)
}
