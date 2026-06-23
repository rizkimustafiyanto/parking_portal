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
		invoice.POST("", middleware.Role("officer"), h.Create)
		invoice.PUT("/:id", middleware.Role("officer"), h.Update)
		invoice.DELETE("/:id", middleware.Role("officer"), h.Delete)
	}

	invoices := router.Group("/invoices")
	invoices.Use(middleware.Auth(jwtSecret))
	{
		invoices.GET("/:id/detail", h.GetDetail)
	}

	members := router.Group("/members")
	members.Use(middleware.Auth(jwtSecret))
	{
		members.GET("/:id/transactions", h.GetMemberTransactions)
		members.GET("/:id/balance-history", h.GetMemberBalanceHistory)
	}
}
