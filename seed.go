package main

import (
	"log"
	"nls-auth/internal/models"
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
