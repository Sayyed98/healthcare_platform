package model

type Hospital struct {
	ID      int64
	Name    string
	Address string
}

type Department struct {
	ID         int64
	HospitalID int64
	Name       string
}

type Doctor struct {
	ID             int64
	HospitalID     int64
	DepartmentID   int64
	Name           string
	Specialization string
}

type Patient struct {
	ID       int64
	Name     string
	Disease  string
	DoctorID int64
}
