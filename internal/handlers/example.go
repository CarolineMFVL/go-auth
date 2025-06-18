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

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    utils.InitLogger()
    database.InitDB()

    if os.Getenv("SEED_DB") == "1" {
        log.Println("Base de données seedée")
        return
    }

}
