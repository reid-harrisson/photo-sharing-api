package server

import (
	"time"

	"github.com/gin-contrib/cors"
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

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Server{
		Router:         gin.Default(),
		SupabaseClient: supabaseClient,
		DB:             postgrestDB,
	}
}

func (server *Server) Start(serverPort string) {
	server.Router.Run(":" + serverPort)
}
