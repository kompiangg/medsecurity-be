package service

import (
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service/auth"
	"medsecurity/service/patient"
	"medsecurity/service/patient_image"
	"medsecurity/service/ping"
)

type Service struct {
	Ping         ping.ServiceItf
	Auth         auth.Service
	Patient      patient.Service
	PatientImage patient_image.Service
}

func New(
	repository repository.Repository,
	config config.Config,
	validator validator.ValidatorItf,
) (Service, error) {
	pingService := ping.New()
	authService := auth.New(
		auth.Config{
			AES: config.AES,
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
	patientImage := patient_image.New(
		repository.Patient,
		repository.Doctor,
		repository.PatientSecret,
		repository.PatientImage,
		repository.Cloudinary,
		repository.AccessHistory,
		validator,
		patient_image.Config{
			AES: config.AES,
		},
	)

	return Service{
		Ping:         pingService,
		Auth:         authService,
		Patient:      patientService,
		PatientImage: patientImage,
	}, nil
}
