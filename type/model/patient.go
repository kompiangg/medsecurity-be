package model

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID                 uuid.UUID `db:"id"`
	DateOfBirth        time.Time `db:"date_of_birth"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
	Password           string    `db:"password"`
	FullName           string    `db:"full_name"`
	BloodType          string    `db:"blood_type"`
	Email              string    `db:"email"`
	Phone              string    `db:"phone"`
	Occupation         string    `db:"occupation"`
	Religion           string    `db:"religion"`
	RelationshipStatus string    `db:"relationship_status"`
	Nationality        string    `db:"nationality"`
	Address            string    `db:"address"`
	Gender             bool      `db:"gender"`
}
