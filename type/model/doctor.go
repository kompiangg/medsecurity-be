package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

type Doctor struct {
	ID           uuid.UUID   `db:"id"`
	PolyclinicID null.String `db:"polyclinic_id"`
	Email        string      `db:"email"`
	Password     string      `db:"password"`
	FullName     string      `db:"full_name"`
	CreatedAt    time.Time   `db:"created_at"`
	UpdatedAt    time.Time   `db:"updated_at"`
}
