package auth

import (
	"context"
	"medsecurity/pkg/validator"
	"medsecurity/repository/auth"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/type/params"
)

type Service interface {
	PatientRegistration(ctx context.Context, param params.ServicePatientRegistrationParam) error
	DoctorRegistration(ctx context.Context, param params.ServiceDoctorRegistrationParam) error
}

type service struct {
	authRepo    auth.Repository
	doctorRepo  doctor.Repository
	patientRepo patient.Repository
	validator   validator.ValidatorItf
}

func New(
	authRepo auth.Repository,
	doctorRepo doctor.Repository,
	patientRepo patient.Repository,
	validator validator.ValidatorItf,
) Service {
	return &service{
		authRepo:    authRepo,
		doctorRepo:  doctorRepo,
		patientRepo: patientRepo,
		validator:   validator,
	}
}
