package routes

import (
	"backend/internal/modules/auth/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}
}
