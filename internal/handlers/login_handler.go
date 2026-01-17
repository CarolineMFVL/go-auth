// Package handlers provides HTTP handler functions for authentication and user management.
//
// This file implements the login handler, JWT token generation and validation, and user lookup by email.
//
// Functions:
//   - LoginHandler: Handles user login, verifies credentials, and returns JWT access and refresh tokens.
//   - ValidateJWT: Validates and decodes JWT tokens, returning claims if valid.
//   - GenerateJWT: Generates signed JWT tokens for access (15 min) or refresh (24h).
//   - GetUserByEmail: Retrieves a user from the database by email.
//
// Types:
//   - Credentials: Represents user login credentials (email and password).
//   - Claims: JWT claims structure including username and registered claims.
//
// Environment Variables:
//   - SECRET_JWT_KEY: Secret key used for signing JWT tokens. Must be set in the environment.
//
// Errors:
//   - Returns appropriate HTTP status codes and error messages for invalid requests, unauthorized access, user not found, and internal server errors.
package handlers

import (
	"log"
	"os"
	"time"

	"nls-auth/internal/handlers/database"
	"nls-auth/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey []byte

func init() {
	secret := os.Getenv("SECRET_JWT_KEY")
	if secret == "" {
		log.Fatal("SECRET_JWT_KEY environment variable is not set or is empty")
	}
	jwtKey = []byte(secret)
}

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
func LoginHandler(c *fiber.Ctx) error {
	var creds Credentials
	err := c.BodyParser((&creds))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	// Find user in database by email
	user, err := GetUserByEmail(creds.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	// Compare password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	accessTokenString, err := GenerateJWT(user.ID, creds.Email, "access")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create access token"})
	}
	refreshTokenString, err := GenerateJWT(user.ID, creds.Email, "refresh")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create refresh token"})
	}

	c.JSON(fiber.Map{
		"refresh_token": refreshTokenString,
		"access_token":  accessTokenString,
	})
	return nil
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
func GenerateJWT(userID uuid.UUID, email string, tokenType string) (string, error) {
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

// Example function to get user by email
func GetUserByEmail(email string) (*models.User, error) {
	/* context := c.Context()
	if context == nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Context is nil")
	} */
	db := database.GetDB() // Assuming you have a GetDB() that returns *gorm.DB
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
