package routes

import (
	"net/http"
	"photo-sharing-api/server"
	"photo-sharing-api/server/handlers"

	docs "photo-sharing-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *server.Server) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiV1 := server.Router.Group("/api/v1")

	groupHealth := apiV1.Group("/health")
	GroupHealth(groupHealth, server)

	groupStorage := apiV1.Group("/storage")
	GroupStorage(groupStorage, server)

	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}

func GroupHealth(group *gin.RouterGroup, server *server.Server) {
	handler := handlers.NewHealthHandler(server)
	group.GET("/", handler.HealthCheck)
}

func GroupStorage(group *gin.RouterGroup, server *server.Server) {
	handler := handlers.NewStorageHandler(server)
	group.POST("/", handler.UploadImage)
}
