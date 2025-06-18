package handlers

import (
    "encoding/json"
    "net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur créé"})
}
