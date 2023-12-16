package auth

import (
	"context"
	"database/sql"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
)

func (s service) DoctorRegistration(ctx context.Context, param params.ServiceDoctorRegistrationParam) error {
	err := s.validator.Validate(param)
	if err != nil {
		return err
	}

	doctor, err := s.doctorRepo.FindDoctorByEmail(ctx, params.RepoFindDoctorByEmailParam{
		Email: param.Email,
	})
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	} else if err != nil {
		return errors.Wrap(err, "error when finding doctor by email")
	}

	if doctor.Email != "" {
		return errors.ErrEmailDuplicated
	}

	err = param.HashPassword()
	if err != nil {
		return errors.Wrap(err, "error when hashing password")
	}

	doctor = param.ToDoctorModel()
	tx, err := s.authRepo.DoctorRegistration(ctx, doctor)
	if err != nil {
		return errors.Wrap(err, "error when registering patient")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "error when commiting transaction")
	}

	return nil
}
