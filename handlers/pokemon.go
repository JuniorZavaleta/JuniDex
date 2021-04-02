package handlers

import (
	"encoding/json"
	"junidex/entities"
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
