package api

import (
	"net/http"

	"log"
	"os"
	"photo-sharing-api/db"
	s "photo-sharing-api/server"
	"photo-sharing-api/server/routes"

	storage_go "github.com/supabase-community/storage-go"
)

var (
	server *s.Server
)

func init() {
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
	if postgresUser == "" || postgresHost == "" || postgresPassword == "" || postgresDatabase == "" || postgresPort == "" {
		log.Println("Failed to load environment.")
		return
	}

	supabaseClient := storage_go.NewClient(supabaseUrl+"/storage/v1", supabaseAnonKey, nil)
	postgresDB := db.Init(postgresUser, postgresHost, postgresPassword, postgresDatabase, postgresPort)

	server := s.NewServer(supabaseClient, postgresDB)

	routes.ConfigureRoutes(server)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Router.ServeHTTP(w, r)
}
