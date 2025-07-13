package handlers

import (
	"net/http"
)

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement token refresh logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token refreshed"))
}
