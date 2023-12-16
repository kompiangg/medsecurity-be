package auth

import (
	"context"
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_secret"
	"medsecurity/type/params"
	"medsecurity/type/result"
)

type Service interface {
	PatientRegistration(ctx context.Context, param params.ServicePatientRegistrationParam) error
	DoctorRegistration(ctx context.Context, param params.ServiceDoctorRegistrationParam) error
	PatientLogin(ctx context.Context, param params.ServicePatientLoginParam) (result.ServicePatientLogin, error)
	DoctorLogin(ctx context.Context, param params.ServiceDoctorLoginParam) (result.ServiceDoctorLogin, error)
}

type service struct {
	config        Config
	doctorRepo    doctor.Repository
	patientRepo   patient.Repository
	patientSecret patient_secret.Repository
	validator     validator.ValidatorItf
}

type Config struct {
	RSA config.RSA
	JWT config.JWTMap
}

func New(
	config Config,
	doctorRepo doctor.Repository,
	patientRepo patient.Repository,
	patientSecret patient_secret.Repository,
	validator validator.ValidatorItf,
) Service {
	return &service{
		config:        config,
		doctorRepo:    doctorRepo,
		patientRepo:   patientRepo,
		patientSecret: patientSecret,
		validator:     validator,
	}
}
