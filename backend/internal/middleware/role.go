package middleware

import (
	"net/http"
	"strings"

	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func Role(required string) gin.HandlerFunc {
	required = strings.TrimSpace(required)

	return func(c *gin.Context) {
		if required == "" {
			c.Next()
			return
		}

		role := strings.TrimSpace(GetRole(c))
		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("unauthorized"))
			return
		}

		if !strings.EqualFold(role, required) {
			c.AbortWithStatusJSON(http.StatusForbidden, response.Error("forbidden"))
			return
		}

		c.Next()
	}
}
