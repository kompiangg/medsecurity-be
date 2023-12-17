package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

type AccessHistory struct {
	ID             uuid.UUID   `db:"id"`
	PatientImageID uuid.UUID   `db:"patient_image_id"`
	PatientID      null.String `db:"patient_id"`
	DoctorID       null.String `db:"doctor_id"`
	Purpose        string      `db:"purpose"`
	CreatedAt      time.Time   `db:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at"`
}
