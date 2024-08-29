package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/GoChatter/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Setting up env vars
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatalf("Error loading .env file: %v\n", envErr)
	}
	port := os.Getenv("LOCALHOST_PORT")

	// Bootstrapping routes and ws configurations
	app := echo.New()
	routes.Init(app)
	app.Start(port)
}
