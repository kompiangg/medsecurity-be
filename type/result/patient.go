package result

import (
	"medsecurity/type/model"
	"time"

	"github.com/google/uuid"
)

type ServicePatientLogin struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}

type ServiceGetAllPatients struct {
	ID          uuid.UUID `json:"id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
}

func (s *ServiceGetAllPatients) FromPatientModel(param model.Patient) {
	if param.Gender {
		s.Gender = "Male"
	} else {
		s.Gender = "Female"
	}

	s.ID = param.ID
	s.DateOfBirth = param.DateOfBirth
	s.FullName = param.FullName
	s.Phone = param.Phone
}

type ServiceGetDetailPatient struct {
	ID                 uuid.UUID `json:"id"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	FullName           string    `json:"full_name"`
	BloodType          string    `json:"blood_type"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	Occupation         string    `json:"occupation"`
	Religion           string    `json:"religion"`
	RelationshipStatus string    `json:"relationship_status"`
	Nationality        string    `json:"nationality"`
	Address            string    `json:"address"`
	Gender             string    `json:"gender"`
}

func (s *ServiceGetDetailPatient) FromPatientModel(param model.Patient) {
	if param.Gender {
		s.Gender = "Male"
	} else {
		s.Gender = "Female"
	}

	s.ID = param.ID
	s.DateOfBirth = param.DateOfBirth
	s.CreatedAt = param.CreatedAt
	s.UpdatedAt = param.UpdatedAt
	s.FullName = param.FullName
	s.BloodType = param.BloodType
	s.Email = param.Email
	s.Phone = param.Phone
	s.Occupation = param.Occupation
	s.Religion = param.Religion
	s.RelationshipStatus = param.RelationshipStatus
	s.Nationality = param.Nationality
	s.Address = param.Address
}
