package model

import (
	"time"

	"github.com/google/uuid"
)

type PatientImage struct {
	ID        uuid.UUID `db:"id"`
	PatientID uuid.UUID `db:"patient_id"`
	DoctorID  uuid.UUID `db:"doctor_id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	URL       string    `db:"url"`
	IsValid   bool      `db:"is_valid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
