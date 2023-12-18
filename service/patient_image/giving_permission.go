package patient_image

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/aesx"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

func (s service) GivingPermission(ctx context.Context, param params.ServiceGivingPermission) error {
	err := s.validator.Validate(param)
	if err != nil {
		return err
	}

	patient, err := s.patientRepository.Find(ctx, params.RepoFindPatient{
		ID: null.NewString(param.PatientID, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return errors.ErrAccountNotFound
	} else if err != nil {
		return errors.Wrap(err, "error when finding patient")
	}

	err = param.CompareHashAndPassword(patient.Password)
	if err != nil {
		return errors.Wrap(err, "error when comparing password")
	}

	_, err = s.doctorRepository.Find(ctx, params.RepoFindDoctor{
		ID: null.NewString(param.DoctorID, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return errors.ErrAccountNotFound
	} else if err != nil {
		return errors.Wrap(err, "error when finding doctor")
	}

	imageID, err := uuid.Parse(param.ImageID)
	if err != nil {
		return errors.Wrap(err, "error when parsing image id")
	}

	_, err = s.patientImageRepository.Find(ctx, params.RepositoryFindPatientImage{
		ID:      imageID,
		IsValid: true,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return errors.ErrRecordNotFound
	} else if err != nil {
		return errors.Wrap(err, "error when finding patient image")
	}

	accessRequest, err := param.ToAccessRequestModel(7)
	if err != nil {
		return errors.Wrap(err, "error when converting to access request model")
	}

	err = s.accessRequestRepository.InsertNewRequest(ctx, accessRequest)
	if err != nil {
		return errors.Wrap(err, "error when inserting new access request")
	}

	encryptedPassword, err := aesx.Encrypt([]byte(s.config.AES.Secret), []byte(param.Password))
	if err != nil {
		return errors.Wrap(err, "error when encrypting password")
	}

	err = s.accessRequestRepository.InsertRequestToRedis(ctx, params.RepositoryInsertRequestToRedis{
		RequestID:       accessRequest.ID.String(),
		KeepAliveInDays: 7,
		Password:        string(encryptedPassword),
	})
	if err != nil {
		return errors.Wrap(err, "error when inserting request to redis")
	}

	return nil
}
