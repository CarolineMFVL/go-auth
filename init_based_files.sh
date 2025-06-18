#!/bin/bash

# filepath: /Users/carolinefauvel/Desktop/go/nls-go-auth/init_root_files.sh

set -e
# @todo: modify the variables below to match your project
APP_NAME="nls-auth"
MODULE_NAME="auth"
PORT=4002
DB_PORT=5432
DB_NAME="auth_db"

echo "Create initial files for $APP_NAME project..."

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