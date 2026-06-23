package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/violation/handler"

	"github.com/gin-gonic/gin"
)

func RegisterViolationType(router *gin.RouterGroup, h *handler.ViolationTypeHandler, jwtSecret string) {
	group := router.Group("/violation-types")
	group.Use(middleware.Auth(jwtSecret))
	{
		group.GET("", h.GetAll)
		group.GET("/:id", h.GetByID)
		group.POST("", middleware.Role("admin"), h.Create)
		group.PUT("/:id", middleware.Role("admin"), h.Update)
		group.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
