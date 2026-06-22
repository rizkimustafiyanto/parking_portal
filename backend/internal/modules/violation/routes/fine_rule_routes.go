package routes

import (
	"backend/internal/middleware"
	"backend/internal/modules/violation/handler"

	"github.com/gin-gonic/gin"
)

func RegisterFineRuleVersion(router *gin.RouterGroup, h *handler.FineRuleVersionHandler, jwtSecret string) {
	group := router.Group("/fine-rule-versions")
	group.Use(middleware.Auth(jwtSecret))
	{
		group.GET("", h.GetAll)
		group.GET("/:id", h.GetByID)
		group.POST("", middleware.Role("admin"), h.Create)
		group.PUT("/:id", middleware.Role("admin"), h.Update)
		group.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}

func RegisterFineRuleDetail(router *gin.RouterGroup, h *handler.FineRuleDetailHandler, jwtSecret string) {
	group := router.Group("/fine-rule-details")
	group.Use(middleware.Auth(jwtSecret))
	{
		group.GET("", h.GetAll)
		group.GET("/:id", h.GetByID)
		group.POST("", middleware.Role("admin"), h.Create)
		group.PUT("/:id", middleware.Role("admin"), h.Update)
		group.DELETE("/:id", middleware.Role("admin"), h.Delete)
	}
}
