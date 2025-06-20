package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("SECRET_JWT_KEY")) // À stocker dans une variable d'env en prod

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// @Summary Login user
// @Description Login and tokens JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {object} map[string]string "JWT Tokens"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Find user in database by email
	// If user does not exist 404 Not Found
	userID := uuid.New().String()
	// Compare password with bcrypt
	if creds.Password != "testpassword" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessTokenString, _ := GenerateJWT(userID, creds.Email, "refresh")
	refreshTokenString, err := GenerateJWT(userID, creds.Email, "refresh")
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"refresh_token": refreshTokenString, "access_token": accessTokenString})
}

// ValidateJWT verifies and decodes a JWT token.
// Returns the claims if the token is valid, otherwise an error.
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// GenerateJWT creates a signed JWT token.
// userID: unique user identifier (string, usually a UUID)
// email: user's email
// tokenType: "access" (15 min) or "refresh" (24h)
func GenerateJWT(userID string, email string, tokenType string) (string, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	if tokenType == "access" {
		exp = time.Now().Add(time.Minute * 15).Unix() // 15 minutes pour un token d'accès
	}
	claims := jwt.MapClaims{
		"user_id":    userID,
		"email":      email,
		"exp":        exp,
		"token_type": tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
