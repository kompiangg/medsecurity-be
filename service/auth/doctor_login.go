package auth

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/type/result"
)

func (s service) DoctorLogin(ctx context.Context, param params.ServiceDoctorLoginParam) (result.ServiceDoctorLogin, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServiceDoctorLogin{}, err
	}

	doctor, err := s.doctorRepo.FindDoctorByEmail(ctx, params.RepoFindDoctorByEmailParam{
		Email: param.Email,
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServiceDoctorLogin{}, errors.ErrAccountNotFound
	} else if err != nil {
		return result.ServiceDoctorLogin{}, errors.Wrap(err, "error at find doctor by email")
	}

	param.ID = doctor.ID

	err = param.ComparePassword(doctor.Password)
	if err != nil {
		return result.ServiceDoctorLogin{}, errors.Wrap(err, "error at compare password")
	}

	res, err := param.GenerateAccessToken(s.config.JWT["doctor"].DurationInDay, s.config.JWT["doctor"].Secret)
	if err != nil {
		return result.ServiceDoctorLogin{}, errors.Wrap(err, "error at generate access token")
	}

	return res, nil
}
