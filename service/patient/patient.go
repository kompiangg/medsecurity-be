package patient

import (
	"context"
	"medsecurity/pkg/validator"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_image"
	"medsecurity/type/params"
	"medsecurity/type/result"
)

type Service interface {
	FindPatientImageBriefInformation(ctx context.Context, param params.ServiceFindPatientImage) ([]result.PatientImageBriefInformation, error)
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
