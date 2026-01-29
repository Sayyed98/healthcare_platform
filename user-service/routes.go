package main

import (
	"hms/user-service/handler"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, h *handler.UserHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		r.GET("/auth/me", h.Me) // ✅ ADD THIS
	}
}
