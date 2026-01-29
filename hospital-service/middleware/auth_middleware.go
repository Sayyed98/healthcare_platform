package middleware

import (
	"log"
	"net/http"

	"hms/hospital-service/grpc_client"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authClient *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		log.Println("SESSION FROM COOKIE:", sessionID, "ERR:", err)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing session"})
			return
		}

		resp, err := authClient.ValidateSession(
			c.Request.Context(),
			sessionID,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		c.Set("user_id", resp.UserId)
		c.Set("role", resp.Role)
		c.Next()

	}
}
