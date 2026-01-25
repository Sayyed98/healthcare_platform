package main

import (
	"hms/hospital-service/handler"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, h *handler.HospitalHandler) {
	r.POST("/hospitals", h.CreateHospital)
	r.POST("/assign-doctor", h.AssignDoctor)
}
