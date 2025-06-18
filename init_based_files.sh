#!/bin/bash

# filepath: /Users/carolinefauvel/Desktop/go/nls-go-auth/init_root_files.sh

set -e
# @todo: modify the variables below to match your project
APP_NAME="nls-auth"
MODULE_NAME="auth"
PORT=4002
DB_PORT=5433
DB_NAME="nls_db"

echo "Create initial files for $APP_NAME project..."

# api/v1/auth/db/postgres.go
cat > api/v1/auth/db/postgres.go <<EOF
package db

import (
	"fmt"
	"log"
	"nls-auth/internal/models"
	"os"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL: ", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
}

type PG_DB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PG_DB {
	return &PG_DB{DB: db}
}
EOF

# constants/application.go
cat > internal/constants/application.go <<EOF
package constants

type AppKey string
type ctxKey string

const (
	ApplicationCtx     ctxKey = "application"
	AuthTokenCtx       ctxKey = "auth_token"
	AuthTokenParsedCtx ctxKey = "auth_token_parsed"
)

type localsKey string

const (
	ConfigLocals    localsKey = "config"
	JwtUserLocals   localsKey = "jwt_user"
	DBLocals        localsKey = "db"
	RequestDBLocals localsKey = "request_db"
)
EOF

# Database configuration
cat > internal/handlers/database/orm.go <<EOF
package database

import (
	"database/sql"
	"fmt"
	"log"
	"nls-auth/internal/models"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func init() {
	_ = godotenv.Load()
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")

	// Connexion temporaire à la base 'postgres'
	connStr := os.Getenv("CONNEXION_STRING")

	// postgresDsn := fmt.Sprintf("host=%s user=%s password=%s dbname=nls_db port=%s sslmode=disable", host, user, password, port)
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL (postgres): ", err)
	}
	defer sqlDB.Close()

	// Vérifie si la base existe, sinon la crée
	var exists bool
	checkQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbname)
	err = sqlDB.QueryRow(checkQuery).Scan(&exists)
	if err == sql.ErrNoRows {
		_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			log.Fatal("Erreur création DB: ", err)
		}
	} else if err != nil && err != sql.ErrNoRows {
		log.Fatal("Erreur vérification DB: ", err)
	}

	// Connexion GORM normale
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erreur connexion PostgreSQL: ", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erreur migration DB: ", err)
	}
	return db, err
}
EOF

cat > internal/handlers/database/constants.go <<EOF
package database

const (
	DefaultDatabase = "gorm"
	DefaultDBUser   = "gorm"
	DefaultHost     = "myhost"
	DefaultPort     = "myport"
	DefaultPassword = "mypassword"
)
EOF

# Utils for logging
cat > internal/utils/logger.go <<EOF
package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if os.Getenv("LOG_JSON") == "1" {
		log.Logger = log.Output(os.Stdout) // JSON
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
}
EOF

# Middlewares
cat > internal/middlewares/request_db_middleware.go <<EOF
package middlewares

import (
	"$APP_NAME/api/v1/auth/db"
	"$APP_NAME/internal/handlers/database/constants"

	"github.com/gofiber/fiber/v2"
)

func RequestDBMiddleware(c *fiber.Ctx) error {
	//application := c.Context().Value(constants.ApplicationCtx).(constants.AppKey)
	DB := c.Locals(constants.DBLocals).(*db.PG_DB)

	if c != nil {
		c.Locals(constants.RequestDBLocals, DB)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return c.Next()
}
EOF

cat > internal/middlewares/jwt_middleware.go <<EOF
package middlewares

import (

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	// Vérification du token JWT
	return c.Next()
}
EOF

# main.go
cat > main.go <<EOF
package main

import (
    "log"
    "net/http"
    "$APP_NAME/internal/handlers"
    "$APP_NAME/internal/handlers/database"
    "os"

    _ "$APP_NAME/docs"
    "$APP_NAME/internal/utils"

    "github.com/gorilla/mux"
    httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
    utils.InitLogger()
    database.InitDB()

    if os.Getenv("SEED_DB") == "1" {
        log.Println("Base de données seedée")
        return
    }

    r := mux.NewRouter()
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
    r.HandleFunc("/example", handlers.ExampleHandler).Methods("POST")
    r.HandleFunc("/ws/{threadId}", handlers.HandleWebSocket)
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bienvenue sur l'API Messaging"))
    })

    log.Println("Serveur sur :$PORT")
    log.Fatal(http.ListenAndServe(":$PORT", r))
}
EOF

# Files
cat > internal/handlers/example.go <<EOF
package main

import (
    "log"
    "net/http"
    "$APP_NAME/internal/handlers"
    "$APP_NAME/internal/handlers/database"
    "os"

    _ "$APP_NAME/docs"
    "$APP_NAME/internal/utils"

    "github.com/gorilla/mux"
    httpSwagger "github.com/swaggo/http-swagger"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    utils.InitLogger()
    database.InitDB()

    if os.Getenv("SEED_DB") == "1" {
        log.Println("Base de données seedée")
        return
    }

}
EOF


# Handler exemple
cat > internal/handlers/login_handler.go <<EOF
package handlers

import (
    "encoding/json"
    "net/http"
)

type Credentials struct {
    Username string \`json:"username"\`
    Password string \`json:"password"\`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    json.NewDecoder(r.Body).Decode(&creds)
    // Authentification fictive
    json.NewEncoder(w).Encode(map[string]string{"token": "demo-token"})
}
EOF

# Handler register
cat > internal/handlers/register_handler.go <<EOF
package handlers

import (
    "encoding/json"
    "net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur créé"})
}
EOF

# Handler websocket
cat > internal/handlers/ws_handler.go <<EOF
package handlers

import (
    "net/http"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("WebSocket endpoint"))
}
EOF


# seed.go
cat > seed.go <<EOF
package main

import (
	"log"
	"$APP_NAME/internal/models"
    "os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDB() {
	var DB *gorm.DB
    username := os.Getenv("FAKE_USER")
    password := os.Getenv("FAKE_PASSWORD")
    email := os.Getenv("FAKE_EMAIL")
	users := []models.User{
		{Username: username, Password: password, email: email},
	}

	for _, u := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Erreur hash %s: %v", u.Username, err)
			continue
		}
		u.Password = string(hashed)
		result := DB.Create(&u)
		if result.Error != nil {
			log.Printf("Erreur insertion %s: %v", u.Username, result.Error)
		} else {
			log.Printf("Utilisateur seedé : %s", u.Username)
		}
	}
}
EOF

# models.go
cat > internal/models/models.go <<EOF
package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string // Hashé (en prod)
	Email    string `gorm:"unique"`
}
EOF


echo "Base files created successfully for $APP_NAME project."