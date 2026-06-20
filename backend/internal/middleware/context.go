package middleware

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) string {
	return c.GetString(ContextUserIDKey)
}

func GetRole(c *gin.Context) string {
	return c.GetString(ContextRoleKey)
}
