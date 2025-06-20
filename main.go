// @title API Authentification
// @version 1.0
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
	r.HandleFunc("/ws/{threadId}", handlers.HandleWebSocket)
	r.HandleFunc("/auth/login", handlers.LoginHandler)
	//r.HandleFunc("/auth/register", handlers.RegisterHandler)
	// r.HandleFunc("/auth/refresh", handlers.RefreshHandler)
	//r.HandleFunc("/auth/logout", handlers.LogoutHandler)
	//r.HandleFunc("/auth/verify", handlers.VerifyHandler)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur l'API Messaging"))
	})

	log.Println("Serveur sur :4002")
	log.Fatal(http.ListenAndServe(":4002", r))
}
