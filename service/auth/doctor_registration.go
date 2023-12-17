package auth

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"

	"github.com/volatiletech/null/v9"
)

func (s service) DoctorRegistration(ctx context.Context, param params.ServiceDoctorRegistrationParam) error {
	err := s.validator.Validate(param)
	if err != nil {
		return err
	}

	doctor, err := s.doctorRepo.Find(ctx, params.RepoFindDoctor{
		Email: null.NewString(param.Email, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
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
	tx, err := s.doctorRepo.Insert(ctx, doctor)
	if err != nil {
		return errors.Wrap(err, "error when registering patient")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "error when commiting transaction")
	}

	return nil
}
