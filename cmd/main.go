package main

import (
	"log"
	"os"
	"photo-sharing-api/db"
	"photo-sharing-api/server"
	"photo-sharing-api/server/routes"

	"github.com/joho/godotenv"
	storage_go "github.com/supabase-community/storage-go"
)

// @Title Photo Sharing API
// @Version 1.0
// @BasePath /api/v1
// @Description RESTful API endpoints for Photo Sharing Application
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
		return
	}

	postgresUser, postgresHost, postgresPassword, postgresDatabase, postgresPort := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT")
	// if postgresUser == "" || postgresHost == "" || postgresPassword == "" || postgresDatabase == "" || postgresPort == "" {
	// 	log.Println("Failed to load environment.")
	// 	return
	// }

	supabaseClient := storage_go.NewClient(supabaseUrl+"/storage/v1", supabaseAnonKey, nil)
	postgresDB := db.Init(postgresUser, postgresHost, postgresPassword, postgresDatabase, postgresPort)

	server := server.NewServer(supabaseClient, postgresDB)

	routes.ConfigureRoutes(server)

	server.Start(serverPort)
}
