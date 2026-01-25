package model

type CreateHospitalRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
}

type AssignDoctorRequest struct {
	PatientName string `json:"patient_name"`
	Disease     string `json:"disease"`
}
