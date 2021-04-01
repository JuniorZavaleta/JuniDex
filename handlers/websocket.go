package handlers

import (
	"fmt"
	"junidex/ws"
	"net/http"

	"github.com/gorilla/mux"
)

type WebsocketHandler struct {
	// can add services
	pool *ws.Pool
}

func NewWebsocketHandler(r *mux.Router, pool *ws.Pool) {
	handler := &WebsocketHandler{pool: pool}

	r.HandleFunc("/ws", handler.serveWs)
}

// Websocket
func (h *WebsocketHandler) serveWs(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WebSocket Endpoint Hit")
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &ws.Client{
		Conn: conn,
		Pool: h.pool,
	}

	h.pool.Register <- client
	client.Read()
}
