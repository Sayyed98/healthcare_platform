package service

import (
	"errors"
	"hms/hospital-service/model"
	"hms/hospital-service/repository"
	"net/http"
)

type HospitalService struct {
	repo *repository.HospitalRepository
}

func NewHospitalService(repo *repository.HospitalRepository) *HospitalService {
	return &HospitalService{repo: repo}
}
func validateUserSession(cookie string) error {
	req, err := http.NewRequest(
		"GET",
		"http://localhost:8080/auth/me",
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unauthorized")
	}

	return nil
}

func (s *HospitalService) CreateHospital(req model.CreateHospitalRequest) (*model.HospitalResponse, error) {
	h := &model.Hospital{
		Name:    req.Name,
		Address: req.Address,
	}

	if err := s.repo.CreateHospital(h); err != nil {
		return nil, err
	}

	return &model.HospitalResponse{
		ID:   h.ID,
		Name: h.Name,
	}, nil
}

func (s *HospitalService) AssignDoctor(req model.AssignDoctorRequest, cookie string) error {

	// 🔐 AUTH CHECK
	if err := validateUserSession(cookie); err != nil {
		return err
	}

	// BUSINESS LOGIC
	doctorID, err := s.repo.FindDoctorByDisease(req.Disease)
	if err != nil {
		return errors.New("no doctor found")
	}

	patient := &model.Patient{
		Name:     req.PatientName,
		Disease:  req.Disease,
		DoctorID: doctorID,
	}

	return s.repo.AssignPatient(patient)
}
