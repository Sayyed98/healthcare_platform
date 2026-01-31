package repository

import (
	"database/sql"
	"hms/hospital-service/model"
)

type HospitalRepository struct {
	db *sql.DB
}

func NewHospitalRepository(db *sql.DB) *HospitalRepository {
	return &HospitalRepository{db: db}
}

func (r *HospitalRepository) CreateHospital(h *model.Hospital) error {
	query := `INSERT INTO hospitals (name, address) VALUES ($1,$2) RETURNING id`
	return r.db.QueryRow(query, h.Name, h.Address).Scan(&h.ID)
}

func (r *HospitalRepository) FindDoctorByDisease(disease string) (int64, error) {
	query := `
		SELECT d.id
		FROM doctors d
		JOIN departments dp ON dp.id = d.department_id
		WHERE dp.name ILIKE $1
		LIMIT 1
	`
	var doctorID int64
	err := r.db.QueryRow(query, "%"+disease+"%").Scan(&doctorID)
	return doctorID, err
}

func (r *HospitalRepository) AssignPatient(p *model.Patient) error {
	query := `
		INSERT INTO patients (name, disease, doctor_id)
		VALUES ($1,$2,$3)
		RETURNING id
	`
	return r.db.QueryRow(query, p.Name, p.Disease, p.DoctorID).Scan(&p.ID)
}
