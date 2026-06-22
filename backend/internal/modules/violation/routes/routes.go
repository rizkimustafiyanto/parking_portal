package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/violation/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler, jwtSecret string) {
	violations := router.Group("/violations")
	violations.Use(middleware.Auth(jwtSecret))
	{
		violations.GET("", h.GetAll)
		violations.GET("/:id", h.GetByID)
		violations.POST("", middleware.Role("admin"), h.Create)
		violations.PUT("/:id", middleware.Role("admin"), h.Update)
		violations.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
