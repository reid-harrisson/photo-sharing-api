package server

import (
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
)

type Server struct {
	Router         *gin.Engine
	SupabaseClient *storage_go.Client
}

func NewServer(supabaseUrl, supabaseKey string) *Server {
	client := storage_go.NewClient(supabaseUrl+"/storage/v1", supabaseKey, nil)

	return &Server{
		Router:         gin.Default(),
		SupabaseClient: client,
	}
}

func (server *Server) Start(serverPort string) {
	server.Router.Run(":" + serverPort)
}
