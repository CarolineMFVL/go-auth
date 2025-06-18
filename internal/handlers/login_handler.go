package handlers

import (
    "encoding/json"
    "net/http"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)
    // Authentification fictive
    json.NewEncoder(w).Encode(map[string]string{"token": "demo-token"})
}
