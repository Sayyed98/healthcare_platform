CREATE TABLE hospitals (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE departments (
    id BIGSERIAL PRIMARY KEY,
    hospital_id BIGINT REFERENCES hospitals(id),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE doctors (
    id BIGSERIAL PRIMARY KEY,
    hospital_id BIGINT REFERENCES hospitals(id),
    department_id BIGINT REFERENCES departments(id),
    name VARCHAR(255) NOT NULL,
    specialization VARCHAR(100)
);

CREATE TABLE patients (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    disease VARCHAR(100),
    doctor_id BIGINT REFERENCES doctors(id)
);
