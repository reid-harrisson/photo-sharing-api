package api

import (
	"net/http"

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
	supabaseUrl, supabaseAnonKey := os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY")

	postgresUser, postgresHost, postgresPassword, postgresDatabase, postgresPort := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT")

	supabaseClient := storage_go.NewClient(supabaseUrl+"/storage/v1", supabaseAnonKey, nil)
	postgresDB := db.Init(postgresUser, postgresHost, postgresPassword, postgresDatabase, postgresPort)

	server := s.NewServer(supabaseClient, postgresDB)

	routes.ConfigureRoutes(server)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server.Router.ServeHTTP(w, r)
}
