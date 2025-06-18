package handlers

import (
    "net/http"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("WebSocket endpoint"))
}
