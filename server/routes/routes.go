package routes

import (
	"gin-test/server"
	"gin-test/server/handlers"

	docs "gin-test/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *server.Server) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiV1 := server.Router.Group("/api/v1")

	groupHealth := apiV1.Group("/health")
	GroupHealth(groupHealth, server)

	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func GroupHealth(group *gin.RouterGroup, server *server.Server) {
	handler := handlers.NewHandlerHealth(server)
	group.GET("/", handler.HealthCheck)
}
