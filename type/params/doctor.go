package params

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/model"
	"medsecurity/type/result"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

type ServiceDoctorLoginParam struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`

	ID uuid.UUID `json:"-"`
}

func (p ServiceDoctorLoginParam) ComparePassword(encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(p.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.ErrIncorrectPassword
	} else if err != nil {
		return err
	}

	return nil
}

func (p ServiceDoctorLoginParam) GenerateAccessToken(day int, secret string) (result.ServiceDoctorLogin, error) {
	var err error
	var res result.ServiceDoctorLogin

	expirationDuration := time.Duration(24*day) * time.Hour

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
