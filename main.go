package main

import (
	"log"
	"net/http"
	"nls-auth/internal/handlers"
	"nls-auth/internal/handlers/database"
	"os"

	_ "nls-auth/docs"
	"nls-auth/internal/utils"

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
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	log.Println("Serveur sur :4001")
	log.Fatal(http.ListenAndServe(":4001", r))
}
