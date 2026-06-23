package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/payment-transaction/handler"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, h *handler.Handler, jwtSecret string) {
	payments := router.Group("/payment")
	payments.Use(middleware.Auth(jwtSecret))
	{
		payments.GET("", h.GetAll)
		payments.GET("/:id", h.GetByID)
		payments.POST("", middleware.Role("admin"), h.Create)
		payments.PUT("/:id", middleware.Role("admin"), h.Update)
		payments.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
