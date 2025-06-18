package handlers

import (
	"log"
	"net/http"
	"nls-auth/internal/handlers/database"
	"os"

	_ "nls-auth/docs"
	"nls-auth/internal/utils"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	utils.InitLogger()
	database.InitDB()

	if os.Getenv("SEED_DB") == "1" {
		log.Println("Base de données seedée")
		return
	}

}
