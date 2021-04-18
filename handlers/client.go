package handlers

import (
	"html/template"
	"junidex/helpers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ClientHandler struct {
	// can add services
}

func NewClientHandler(r *mux.Router) {
	handler := &ClientHandler{}

	r.HandleFunc("/", handler.IndexView).Methods("GET")
}

func (h *ClientHandler) IndexView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(helpers.GetTemplateFilepath("client", "index.html"))

	if err != nil {
		log.Panic(err)
	}

	t.Execute(w, nil)
}
