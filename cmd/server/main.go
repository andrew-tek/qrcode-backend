package main

import (
    "log"
    "net/http"

    "qrcode-backend/internal/routes"
)

func main() {
    r := routes.NewRouter() // Use NewRouter to initialize the router

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}