package api

import (
	"net/http"

	"log"
	"os"
	s "photo-sharing-api/server"
	"photo-sharing-api/server/routes"

	"github.com/joho/godotenv"
)

var (
	server *s.Server
)

func init() {
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

	server = s.NewServer(supabaseUrl, supabaseAnonKey)

	routes.ConfigureRoutes(server)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Router.ServeHTTP(w, r)
}
