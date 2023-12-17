package result

import (
	"medsecurity/type/model"
	"time"

	"github.com/google/uuid"
)

type PatientImageBriefInformation struct {
	ID          uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PatientName string    `json:"patient_name"`
	DoctorID    uuid.UUID `json:"doctor_id"`
	DoctorName  string    `json:"doctor_name"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	IsValid     bool      `json:"is_valid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *PatientImageBriefInformation) FromPatientModel(
	model model.PatientImage,
	patient model.Patient,
	doctor model.Doctor,
) {
	p.ID = model.ID
	p.PatientID = model.PatientID
	p.PatientName = patient.FullName
	p.DoctorID = model.DoctorID
	p.DoctorName = doctor.FullName
	p.Name = model.Name
	p.Type = model.Type
	p.IsValid = model.IsValid
	p.CreatedAt = model.CreatedAt
	p.UpdatedAt = model.UpdatedAt
}
