package main

import (
	"hms/user-service/handler"
	"hms/user-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func registerRoutes(r *gin.Engine, h *handler.UserHandler, redis *redis.Client) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.GET("/me", h.Me) // ✅ ADD THIS
		auth.POST("/logout", middleware.AuthMiddleware(redis), h.Logout)
	}
}
