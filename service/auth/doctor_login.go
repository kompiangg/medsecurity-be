package auth

import (
	"context"
	"medsecurity/config"
	"medsecurity/pkg/errors"
	"medsecurity/type/constant"
	"medsecurity/type/params"
	"medsecurity/type/result"

	"github.com/volatiletech/null/v9"
)

func (s service) DoctorLogin(ctx context.Context, param params.ServiceDoctorLoginParam) (result.ServiceDoctorLogin, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServiceDoctorLogin{}, err
	}

	doctor, err := s.doctorRepo.Find(ctx, params.RepoFindDoctor{
		Email: null.NewString(param.Email, true),
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

	res, err := param.GenerateAccessToken(s.config.JWT[config.DoctorJWT].DurationInDay, s.config.JWT[config.DoctorJWT].Secret)
	if err != nil {
		return result.ServiceDoctorLogin{}, errors.Wrap(err, "error at generate access token")
	}

	res.Role = constant.DoctorRole

	return res, nil
}
