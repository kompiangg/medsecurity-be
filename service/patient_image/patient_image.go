package patient_image

import (
	"context"
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository/access_history"
	"medsecurity/repository/access_request"
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
	PatientRequestGetImage(ctx context.Context, param params.ServicePatientRequestGetImage) (result.ServicePatientRequestGetImage, error)
	PatientGetImage(ctx context.Context, param params.ServicePatientGetImage) (result.ServicePatientGetImage, error)
	GivingPermission(ctx context.Context, param params.ServiceGivingPermission) error
	DoctorGetImage(ctx context.Context, param params.DoctorPatientGetImage) (result.ServiceDoctorGetImage, error)
}

type service struct {
	patientRepository       patient.Repository
	doctorRepository        doctor.Repository
	patientSecretRepository patient_secret.Repository
	patientImageRepository  patient_image.Repository
	cloudinaryRepository    cloudinary.Repository
	accessHistoryRepository access_history.Repository
	accessRequestRepository access_request.Repository
	validator               validator.ValidatorItf
	config                  Config
}

type Config struct {
	AES config.AES
}

func New(
	patientRepository patient.Repository,
	doctorRepository doctor.Repository,
	patientSecretRepository patient_secret.Repository,
	patientImageRepository patient_image.Repository,
	cloudinaryRepository cloudinary.Repository,
	accessHistoryRepository access_history.Repository,
	accessRequestRepository access_request.Repository,
	validator validator.ValidatorItf,
	config Config,
) Service {
	return &service{
		patientRepository:       patientRepository,
		doctorRepository:        doctorRepository,
		patientSecretRepository: patientSecretRepository,
		patientImageRepository:  patientImageRepository,
		cloudinaryRepository:    cloudinaryRepository,
		accessHistoryRepository: accessHistoryRepository,
		accessRequestRepository: accessRequestRepository,
		validator:               validator,
		config:                  config,
	}
}
