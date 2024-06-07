package main

import (
	"gin-test/server"
	"gin-test/server/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @Title Gin Test
// @Version 1.0
// @BasePath /api/v1

func main() {
	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load environment.")
		return
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Println("Failed to load environment.")
	}

	server := server.NewServer()

	routes.ConfigureRoutes(server)

	server.Start(serverPort)
}
