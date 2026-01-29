package handler

import (
	"fmt"
	"hms/hospital-service/model"
	"hms/hospital-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HospitalHandler struct {
	service *service.HospitalService
}

func NewHospitalHandler(service *service.HospitalService) *HospitalHandler {
	return &HospitalHandler{service: service}
}

func (h *HospitalHandler) CreateHospital(c *gin.Context) {
	var req model.CreateHospitalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateHospital(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *HospitalHandler) AssignDoctor(c *gin.Context) {
	var req model.AssignDoctorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cookie := c.Request.Header.Get("Cookie")
	fmt.Println("cokkie", cookie)

	if err := h.service.AssignDoctor(req, cookie); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "doctor assigned"})
}
