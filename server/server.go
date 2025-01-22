package server

import (
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
	"gorm.io/gorm"
)

type Server struct {
	Router         *gin.Engine
	SupabaseClient *storage_go.Client
	DB             *gorm.DB
}

func NewServer(supabaseClient *storage_go.Client, postgrestDB *gorm.DB) *Server {
	return &Server{
		Router:         gin.Default(),
		SupabaseClient: supabaseClient,
		DB:             postgrestDB,
	}
}

func (server *Server) Start(serverPort string) {
	server.Router.Run(":" + serverPort)
}
