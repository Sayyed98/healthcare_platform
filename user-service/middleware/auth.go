package middleware

import (
	"net/http"

	"hms/user-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing session",
			})
			return
		}

		userID, err := utils.GetSession(rdb, cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid session",
			})
			return
		}

		// attach user info to context
		c.Set("user_id", userID)

		c.Next()
	}
}
func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}
