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
		payments.POST("", h.Create)
		payments.PUT("/:id", middleware.Role("officer"), h.Update)
		payments.DELETE("/:id", middleware.Role("officer"), h.Delete)
	}
}
