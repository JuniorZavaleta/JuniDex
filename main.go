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

func main() {
	// Checking directory for html templates
	fmt.Printf("DIR TEMPLATE: %s\n", os.Getenv("DIR_TEMPLATE"))
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
	// Adming Client Html handler
	handlers.NewClientHandler(r)

	// Run server
	http.ListenAndServe(":8080", r)
}
