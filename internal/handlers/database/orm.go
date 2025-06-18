package database

import (
	"database/sql"
	"fmt"
	"log"
	"nls-auth/internal/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	_ = godotenv.Load()
}

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
