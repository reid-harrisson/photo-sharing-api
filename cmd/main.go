package main

import (
	"log"
	"os"
	"photo-sharing-api/server"
	"photo-sharing-api/server/routes"

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

	supabaseUrl, supabaseAnonKey := os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY")
	if supabaseUrl == "" || supabaseAnonKey == "" {
		log.Println("Failed to load environment.")
	}

	server := server.NewServer(supabaseUrl, supabaseAnonKey)

	routes.ConfigureRoutes(server)

	server.Start(serverPort)
}
