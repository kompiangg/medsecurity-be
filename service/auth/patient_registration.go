package auth

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
)

func (s service) PatientRegistration(ctx context.Context, param params.ServicePatientRegistrationParam) error {
	err := s.validator.Validate(param)
	if err != nil {
		return err
	}

	patient, err := s.patientRepo.FindPatientByEmail(ctx, params.RepoFindPatientByEmailParam{
		Email: param.Email,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		err = nil
	} else if err != nil {
		return errors.Wrap(err, "error when finding patient by email")
	}

	if patient.Email != "" {
		return errors.ErrEmailDuplicated
	}

	err = param.HashPassword()
	if err != nil {
		return errors.Wrap(err, "error when hashing password")
	}

	patient, err = param.ToPatientModel()
	if err != nil {
		return errors.Wrap(err, "error when converting param to patient model")
	}

	patientTx, err := s.patientRepo.Insert(ctx, patient)
	if err != nil {
		return errors.Wrap(err, "error when registering patient")
	}

	err = patientTx.Commit()
	if err != nil {
		return errors.Wrap(err, "error when commiting patient transaction")
	}

	patientSecret, err := param.ToPatientSecretModel(patient.ID, s.config.RSA.KeySize)
	if err != nil {
		patientTx.Rollback()
		return errors.Wrap(err, "error when converting param to patient secret model")
	}

	patientSecretTx, err := s.patientSecret.Insert(ctx, patientSecret)
	if err != nil {
		patientTx.Rollback()
		return errors.Wrap(err, "error when inserting patient secret")
	}

	err = patientSecretTx.Commit()
	if err != nil {
		patientTx, err := s.patientRepo.DeleteByID(ctx, patient.ID)
		if err != nil {
			return errors.Wrap(err, "error when deleting patient")
		}

		err = patientTx.Commit()
		if err != nil {
			return errors.Wrap(err, "error when commiting patient transaction")
		}

		return errors.Wrap(err, "error when commiting patient secret transaction")
	}

	return nil
}
