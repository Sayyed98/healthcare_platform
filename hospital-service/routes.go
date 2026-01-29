package main

import (
	"hms/hospital-service/handler"

	"github.com/gin-gonic/gin"
)

//	func registerRoutes(r *gin.Engine, h *handler.HospitalHandler, group *gin.RouterGroup) {
//		r.POST("/hospitals", h.CreateHospital)
//		{
//			group.POST("/assign-doctor", h.AssignDoctor)
//		}
//	}
func registerRoutes(r *gin.Engine, h *handler.HospitalHandler, auth gin.HandlerFunc) {

	r.POST("/hospitals", h.CreateHospital)

	protected := r.Group("/")
	protected.Use(auth)
	protected.POST("/assign-doctor", h.AssignDoctor)
}
