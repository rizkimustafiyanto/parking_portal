package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/invoice/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler, jwtSecret string) {
	invoice := router.Group("/invoice")
	invoice.Use(middleware.Auth(jwtSecret))
	{
		invoice.GET("", h.GetAll)
		invoice.GET("/:id", h.GetByID)
		invoice.POST("", middleware.Role("admin"), h.Create)
		invoice.PUT("/:id", middleware.Role("admin"), h.Update)
		invoice.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
