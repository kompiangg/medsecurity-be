package params

import (
	"medsecurity/type/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RepoFindDoctorByEmailParam struct {
	Email string `db:"email"`
}

type ServiceDoctorRegistrationParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

func (p *ServiceDoctorRegistrationParam) HashPassword() error {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Password = string(bcryptPassword)
	return nil
}

func (p ServiceDoctorRegistrationParam) ToDoctorModel() model.Doctor {
	doctor := model.Doctor{
		ID:       uuid.New(),
		Email:    p.Email,
		Password: p.Password,
		FullName: p.FullName,
	}

	return doctor
}
