package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type StaticHandler struct {
	// can add services
}

func NewStaticHandler(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(os.Getenv("STATIC_FOLDER")))))
}
