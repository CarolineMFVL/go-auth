// @title API Authentification
// @version 1.0
package main

import (
	"log"
	"net/http"
	"os"

	"nls-auth/internal/handlers"
	"nls-auth/internal/handlers/database"

	_ "nls-auth/docs"
	"nls-auth/internal/utils"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	utils.InitLogger()
	database.InitDB()

	if os.Getenv("SEED_DB") == "1" {
		log.Println("Base de données seedée")
		return
	}
	// c := new(fiber.Ctx)
	// Initialize Fiber app
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenue sur l'API Authentification")
	})
	app.Use("/swagger", filesystem.New(filesystem.Config{
		Root:   http.Dir("./docs"), // Adjust path if needed
		Browse: true,
		Index:  "index.html",
	}))

	// r := mux.NewRouter()
	// app.Use(middlewares.RequestDBMiddleware(c))
	app.Get("/swagger/*", adaptor.HTTPHandler(httpSwagger.WrapHandler))
	app.Post("/auth/login", handlers.LoginHandler)
	app.Post("/auth/register", handlers.RegisterHandler)
	app.Post("/auth/refresh", handlers.RefreshHandler)
	app.Post("/auth/logout", handlers.LogoutHandler)
	app.Get("/auth/verify", handlers.VerifyHandler)
	/* app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Bienvenue sur l'API Authentification")
		return nil
	}) */

	log.Println("Serveur sur :4002")
	log.Fatal(app.Listen(":4002"))
}
