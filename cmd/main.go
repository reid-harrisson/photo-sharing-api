package main

import (
	"gin-test/server"
	"gin-test/server/routes"
)

// @Title Gin Test
// @Version 1.0
// @BasePath /api/v1

func main() {
	server := server.NewServer()

	routes.ConfigureRoutes(server)

	server.Start("8000")
}
