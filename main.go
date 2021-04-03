package main

import (
	"fmt"
	"junidex/handlers"
	"junidex/ws"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Health Check endpoint
func HandlerPing(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("pong"))
}

const version = "0.2.1"

func main() {
	// Checking directory for html templates
	fmt.Printf("Version: \t%s\n", version)
	fmt.Printf("DIR TEMPLATE: %s\n", os.Getenv("DIR_TEMPLATE"))
	fmt.Printf("STATIC FOLDER: %s\n", os.Getenv("STATIC_FOLDER"))
	fmt.Printf("DATA FOLDER: %s\n", os.Getenv("DATA_FOLDER"))

	// Starting pool
	pool := ws.NewPool()
	go pool.Start()

	r := mux.NewRouter()
	// Adding HealthCheck handler
	r.HandleFunc("/ping", HandlerPing)
	// Adding REST handlers
	handlers.NewRestHandler(r)
	// Adding Websocekt handler
	handlers.NewWebsocketHandler(r, pool)

	// Adding Admin Html handler
	handlers.NewAdminHandler(r)
	// Adding Client Html handler
	handlers.NewClientHandler(r)
	// Adding Static folder handler
	handlers.NewStaticHandler(r)

	// Run server
	http.ListenAndServe(":8080", r)
}
