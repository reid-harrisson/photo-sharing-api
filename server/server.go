package server

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (server *Server) Start(serverPort string) {
	server.Router.Run(":" + serverPort)
}
