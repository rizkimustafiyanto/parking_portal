package routes

import (
	"net/http"

	authhandler "backend/internal/modules/auth/handler"
	authroutes "backend/internal/modules/auth/routes"
	authsvc "backend/internal/modules/auth/service"
	"backend/internal/modules/user/handler"
	userrepo "backend/internal/modules/user/repository"
	userroutes "backend/internal/modules/user/routes"
	usersvc "backend/internal/modules/user/service"
	"backend/pkg/response"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	v1 := router.Group("/api")

	userRepository := userrepo.NewRepository(db)
	userService := usersvc.NewService(userRepository)
	userHandler := handler.NewHandler(userService)
	userroutes.Register(v1, userHandler, jwtSecret)

	authService := authsvc.NewService(userRepository, jwtSecret)
	authHandler := authhandler.NewHandler(authService)
	authroutes.Register(v1, authHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Success("server is running", nil))
	})
}
