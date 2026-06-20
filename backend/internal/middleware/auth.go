package middleware

import (
	"net/http"
	"strings"

	"backend/pkg/jwt"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "user_id"
	ContextRoleKey   = "role"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("missing authorization token"))
			return
		}

		tokenString := authHeader
		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			tokenString = strings.TrimSpace(authHeader[7:])
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("missing authorization token"))
			return
		}

		claims, err := jwt.ValidateToken(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid token"))
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextRoleKey, claims.Role)
		c.Next()
	}
}
