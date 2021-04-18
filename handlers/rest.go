package handlers

import (
	"encoding/json"
	"junidex/entities"
	"junidex/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpRestHandler struct {
	// can add services
}

func NewRestHandler(r *mux.Router) {
	handler := &HttpRestHandler{}

	r.HandleFunc("/pokemon/{id}", handler.GetPokemon).Methods("GET")
	r.HandleFunc("/api/team", handler.GetTeam).Methods("GET")
	r.HandleFunc("/api/all", handler.GetAll).Methods("GET")
}

func (h *HttpRestHandler) GetPokemon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)

	log.Printf("Pokemon ID requested %s", vars["id"])
	bulbasaur := entities.Pokemon{}
	bulbasaur.Name = "Bulbasaur"
	bulbasaur.TypeOne = "Grass"
	bulbasaur.TypeTwo = "Poison"

	json.NewEncoder(w).Encode(bulbasaur)
}

func (h *HttpRestHandler) GetTeam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Printf("Loading Team...\n");

	team := repo.LoadTeam()
	log.Printf("Team loaded: %v", team)

	json.NewEncoder(w).Encode(team)
}

func (h *HttpRestHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	log.Printf("Loading all...\n");

	all := repo.LoadAll()

	json.NewEncoder(w).Encode(all)
}
