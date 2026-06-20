package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/user/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler, jwtSecret string) {
	users := router.Group("/users")
	users.Use(middleware.Auth(jwtSecret))
	{
		users.GET("", h.GetAll)
		users.GET("/:id", h.GetByID)
		users.POST("", middleware.Role("admin"), h.Create)
		users.PUT("/:id", middleware.Role("admin"), h.Update)
		users.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
