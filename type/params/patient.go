package params

import (
	"crypto/rand"
	"crypto/rsa"
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/result"
	"medsecurity/utils/rsax"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
	"golang.org/x/crypto/bcrypt"
)

type RepoFindPatient struct {
	Email null.String `db:"email"`
	ID    null.String `db:"id"`
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

type ServicePatientLoginParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`

	ID uuid.UUID `json:"-"`
}

func (p ServicePatientLoginParam) ComparePassword(encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(p.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.ErrIncorrectPassword
	} else if err != nil {
		return err
	}

	return nil
}

func (p ServicePatientLoginParam) GenerateAccessToken(day int, secret string) (result.ServicePatientLogin, error) {
	var err error
	var res result.ServicePatientLogin

	expirationDuration := time.Duration(24*time.Hour) * time.Duration(day)

	jwtClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
		Subject:   p.ID.String(),
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	res.AccessToken, err = unsignedToken.SignedString([]byte(secret))
	if err != nil {
		return res, errors.Wrap(err, "[Entity] UserAccount (GenerateAccessToken): error on creating jwt access token")
	}

	return res, nil
}
