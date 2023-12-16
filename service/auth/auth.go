package auth

import (
	"context"
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository/auth"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_secret"
	"medsecurity/type/params"
)

type Service interface {
	PatientRegistration(ctx context.Context, param params.ServicePatientRegistrationParam) error
	DoctorRegistration(ctx context.Context, param params.ServiceDoctorRegistrationParam) error
}

type service struct {
	config        Config
	authRepo      auth.Repository
	doctorRepo    doctor.Repository
	patientRepo   patient.Repository
	patientSecret patient_secret.Repository
	validator     validator.ValidatorItf
}

type Config struct {
	RSA config.RSA
}

func New(
	config Config,
	authRepo auth.Repository,
	doctorRepo doctor.Repository,
	patientRepo patient.Repository,
	patientSecret patient_secret.Repository,
	validator validator.ValidatorItf,
) Service {
	return &service{
		config:        config,
		authRepo:      authRepo,
		doctorRepo:    doctorRepo,
		patientRepo:   patientRepo,
		patientSecret: patientSecret,
		validator:     validator,
	}
}
