package service

import (
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service/auth"
	"medsecurity/service/patient"
	"medsecurity/service/ping"
)

type Service struct {
	Ping    ping.ServiceItf
	Auth    auth.Service
	Patient patient.Service
}

func New(
	repository repository.Repository,
	config config.Config,
	validator validator.ValidatorItf,
) (Service, error) {
	pingService := ping.New()
	authService := auth.New(
		auth.Config{
			RSA: config.RSA,
			JWT: config.JWT,
		},
		repository.Doctor,
		repository.Patient,
		repository.PatientSecret,
		validator,
	)
	patientService := patient.New(
		repository.Patient,
		repository.Doctor,
		repository.PatientImage,
		validator,
	)

	return Service{
		Ping:    pingService,
		Auth:    authService,
		Patient: patientService,
	}, nil
}
