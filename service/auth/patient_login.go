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

func (s service) PatientLogin(ctx context.Context, param params.ServicePatientLoginParam) (result.ServicePatientLogin, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServicePatientLogin{}, err
	}

	patient, err := s.patientRepo.Find(ctx, params.RepoFindPatient{
		Email: null.NewString(param.Email, true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServicePatientLogin{}, errors.ErrAccountNotFound
	} else if err != nil {
		return result.ServicePatientLogin{}, errors.Wrap(err, "error at find patient by email")
	}

	param.ID = patient.ID

	err = param.ComparePassword(patient.Password)
	if errors.Is(err, errors.ErrIncorrectPassword) {
		return result.ServicePatientLogin{}, err
	} else if err != nil {
		return result.ServicePatientLogin{}, errors.Wrap(err, "error at compare password")
	}

	res, err := param.GenerateAccessToken(s.config.JWT[config.PatientJWT].DurationInDay, s.config.JWT[config.PatientJWT].Secret)
	if err != nil {
		return result.ServicePatientLogin{}, errors.Wrap(err, "error at generate access token")
	}

	res.Role = constant.PatientRole

	return res, nil
}
