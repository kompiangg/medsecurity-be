package auth

import (
	"context"
	"database/sql"
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
	if errors.Is(err, sql.ErrNoRows) {
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

	err = s.authRepo.PatientRegistration(ctx, patient)
	if err != nil {
		return errors.Wrap(err, "error when registering patient")
	}

	return nil
}
