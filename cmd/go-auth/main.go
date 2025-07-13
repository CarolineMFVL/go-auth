// @title API Authentification
// @version 1.0
package main

import (
	"log"
	"os"

	"nls-auth/internal/handlers"
	"nls-auth/internal/handlers/database"
	"nls-auth/internal/middlewares"

	_ "nls-auth/docs"
	"nls-auth/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InitLogger()
	database.InitDB()

	if os.Getenv("SEED_DB") == "1" {
		log.Println("Base de données seedée")
		return
	}
	c := new(fiber.Ctx)
	// Initialize Fiber app
	app := fiber.New()

	// r := mux.NewRouter()
	app.Use(middlewares.RequestDBMiddleware(c))
	/* app.Get("/swagger/*", adaptator.HTTPHandler(httpSwagger.WrapHandler))
	app.Get("/swagger/doc.json", fiber.HTTPHandler(httpSwagger.Handler().DocJSONHandler())) */
	app.Post("/auth/login", handlers.LoginHandler)
	/* 	app.Post("/auth/register", handlers.RegisterHandler)
	   	app.Post("/auth/refresh", handlers.RefreshHandler)
	   	app.Post("/auth/logout", handlers.LogoutHandler)
	   	app.Get("/auth/verify", handlers.VerifyHandler)
	   	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
	   		w.Write([]byte("Bienvenue sur l'API Messaging"))
	   	}) */

	log.Println("Serveur sur :4002")
	log.Fatal(app.Listen(":4002"))
}
