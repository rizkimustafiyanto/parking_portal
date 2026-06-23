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
		users.POST("", middleware.Role("officer"), h.Create)
		users.PUT("/:id", middleware.Role("officer"), h.Update)
		users.POST("/:id/top-up-balance", middleware.Role("officer"), h.TopUpBalance)
		users.DELETE("/:id", middleware.Role("officer"), h.Delete)
	}
}
