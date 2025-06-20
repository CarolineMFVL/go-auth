package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler_Success(t *testing.T) {
	payload := Credentials{
		Email:    "testuser@example.com",
		Password: "testpassword",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if _, ok := resp["access_token"]; !ok {
		t.Errorf("access_token not found in response")
	}
	if _, ok := resp["refresh_token"]; !ok {
		t.Errorf("refresh_token not found in response")
	}
}

func TestLoginHandler_Unauthorized(t *testing.T) {
	payload := Credentials{
		Email:    "wronguser@example.com",
		Password: "wrongpassword",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr.Code)
	}
}

func TestLoginHandler_BadRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}

func TestLoginHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}
