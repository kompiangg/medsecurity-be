package model

import (
	"time"

	"github.com/google/uuid"
)

type PatientSecret struct {
	ID         uuid.UUID `db:"id"`
	PatientID  uuid.UUID `db:"patient_id"`
	PrivateKey string    `db:"private_key"`
	KeySize    int       `db:"key_size"`
	IsValid    bool      `db:"is_valid"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
