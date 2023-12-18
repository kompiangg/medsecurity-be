package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

type AccessRequest struct {
	ID           uuid.UUID `db:"id"`
	PatientID    uuid.UUID `db:"patient_id"`
	DoctorID     uuid.UUID `db:"doctor_id"`
	ImageID      uuid.UUID `db:"image_id"`
	Purpose      string    `db:"purpose"`
	IsAllowed    null.Bool `db:"is_allowed"`
	AllowedUntil time.Time `db:"allowed_until"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
