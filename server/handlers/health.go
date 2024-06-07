package handlers

import (
	"gin-test/responses"
	"gin-test/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerHealth struct {
	Server *server.Server
}

func NewHandlerHealth(server *server.Server) *HandlerHealth {
	return &HandlerHealth{
		Server: server,
	}
}

// PingExample godoc
// @Summary Health check
// @Schemes
// @Description Health check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Router /health [get]
func (handler *HandlerHealth) HealthCheck(context *gin.Context) {
	responses.MessageResponse(context, http.StatusOK, "Server is running!")
}
