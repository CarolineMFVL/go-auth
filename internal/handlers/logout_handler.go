package handlers

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logout logic (e.g., invalidate tokens, clear cookies)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out"))
}
