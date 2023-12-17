package patient_image

import (
	"context"
	"medsecurity/pkg/validator"
	"medsecurity/repository/access_history"
	"medsecurity/repository/cloudinary"
	"medsecurity/repository/doctor"
	"medsecurity/repository/patient"
	"medsecurity/repository/patient_image"
	"medsecurity/repository/patient_secret"
	"medsecurity/type/params"
	"medsecurity/type/result"
)

type Service interface {
	Insert(ctx context.Context, param params.ServiceCreatePatientImage) error
	FindBriefInformation(ctx context.Context, param params.ServiceFindPatientImage) ([]result.PatientImageBriefInformation, error)
}

type service struct {
	patientRepository       patient.Repository
	doctorRepository        doctor.Repository
	patientSecretRepository patient_secret.Repository
	patientImageRepository  patient_image.Repository
	cloudinaryRepository    cloudinary.Repository
	accessHistoryRepository access_history.Repository
	validator               validator.ValidatorItf
}

func New(
	patientRepository patient.Repository,
	doctorRepository doctor.Repository,
	patientSecretRepository patient_secret.Repository,
	patientImageRepository patient_image.Repository,
	cloudinaryRepository cloudinary.Repository,
	accessHistoryRepository access_history.Repository,
	validator validator.ValidatorItf,
) Service {
	return &service{
		patientRepository:       patientRepository,
		doctorRepository:        doctorRepository,
		patientSecretRepository: patientSecretRepository,
		patientImageRepository:  patientImageRepository,
		cloudinaryRepository:    cloudinaryRepository,
		accessHistoryRepository: accessHistoryRepository,
		validator:               validator,
	}
}
