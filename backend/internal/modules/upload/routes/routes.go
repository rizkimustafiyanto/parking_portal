package routes

import (
	"backend/internal/modules/upload/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler) {
	uploads := router.Group("/uploads")
	{
		uploads.POST("", h.Upload)
	}
}

