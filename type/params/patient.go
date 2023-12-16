package params

import (
	"crypto/rand"
	"crypto/rsa"
	"medsecurity/type/model"
	"medsecurity/utils/rsax"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RepoFindPatientByEmailParam struct {
	Email string `json:"email"`
}

type ServicePatientRegistrationParam struct {
	DateOfBirth        string `json:"date_of_birth" validate:"required"`
	Password           string `json:"password" validate:"required"`
	FullName           string `json:"full_name" validate:"required"`
	BloodType          string `json:"blood_type" validate:"required"`
	Email              string `json:"email" validate:"required"`
	Phone              string `json:"phone" validate:"required"`
	Occupation         string `json:"occupation" validate:"required"`
	Religion           string `json:"religion" validate:"required"`
	RelationshipStatus string `json:"relationship_status" validate:"required"`
	Nationality        string `json:"nationality" validate:"required"`
	Address            string `json:"address" validate:"required"`
	Gender             bool   `json:"gender"`

	UnencryptedPassword string `json:"-" validate:"-"`
}

func (p *ServicePatientRegistrationParam) HashPassword() error {
	p.UnencryptedPassword = p.Password

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Password = string(bcryptPassword)
	return nil
}

func (p ServicePatientRegistrationParam) ToPatientModel() (model.Patient, error) {
	dateOfBirth, err := time.Parse("2006-01-02", p.DateOfBirth)
	if err != nil {
		return model.Patient{}, err
	}

	patient := model.Patient{
		ID:                 uuid.New(),
		DateOfBirth:        dateOfBirth,
		Password:           p.Password,
		FullName:           p.FullName,
		BloodType:          p.BloodType,
		Email:              p.Email,
		Phone:              p.Phone,
		Occupation:         p.Occupation,
		Religion:           p.Religion,
		RelationshipStatus: p.RelationshipStatus,
		Nationality:        p.Nationality,
		Address:            p.Address,
		Gender:             p.Gender,
	}

	return patient, nil
}

func (p ServicePatientRegistrationParam) ToPatientSecretModel(patientID uuid.UUID, keySize int) (model.PatientSecret, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return model.PatientSecret{}, err
	}

	encryptedPrivateKey, err := rsax.EncryptPrivateKey(privateKey, p.UnencryptedPassword)
	if err != nil {
		return model.PatientSecret{}, err
	}

	return model.PatientSecret{
		ID:         uuid.New(),
		PatientID:  patientID,
		PrivateKey: string(encryptedPrivateKey),
		KeySize:    keySize,
		IsValid:    true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
