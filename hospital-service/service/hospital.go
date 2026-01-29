package service

import (
	"errors"
	"hms/hospital-service/grpc_client"
	"hms/hospital-service/model"
	"hms/hospital-service/repository"
)

type HospitalService struct {
	repo       *repository.HospitalRepository
	authClient *grpc_client.AuthClient
}

func NewHospitalService(repo *repository.HospitalRepository, authClient *grpc_client.AuthClient) *HospitalService {
	return &HospitalService{repo: repo, authClient: authClient}
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
