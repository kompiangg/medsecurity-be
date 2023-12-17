package patient

import (
	"medsecurity/pkg/validator"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_image"
)

type Service interface {
}

type service struct {
	patientRepository      patient.Repository
	patientImageRepository patient_image.Repository
	doctorRepository       doctor.Repository
	validator              validator.ValidatorItf
}

func New(
	patientRepository patient.Repository,
	doctorRepository doctor.Repository,
	patientImageRepository patient_image.Repository,
	validator validator.ValidatorItf,
) Service {
	return &service{
		patientRepository:      patientRepository,
		doctorRepository:       doctorRepository,
		patientImageRepository: patientImageRepository,
		validator:              validator,
	}
}
