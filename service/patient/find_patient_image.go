package patient

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/constant"
	"medsecurity/type/params"
	"medsecurity/type/result"

	"github.com/volatiletech/null/v9"
)

func (s service) FindPatientImageBriefInformation(ctx context.Context, param params.ServiceFindPatientImage) ([]result.PatientImageBriefInformation, error) {
	if param.PatientID.Valid {
		_, err := s.patientRepository.Find(ctx, params.RepoFindPatient{
			ID: param.PatientID,
		})
		if errors.Is(err, errors.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		} else if err != nil {
			return nil, errors.Wrap(err, "error when finding patient")
		}

		if param.Role == constant.PatientRole {
			if !param.PatientID.Valid {
				return nil, errors.ErrUnauthorized
			}

			if param.PatientID.String != param.AccountID {
				return nil, errors.ErrUnauthorized
			}
		}
	}

	if param.DoctorID.Valid {
		_, err := s.doctorRepository.Find(ctx, params.RepoFindDoctor{
			ID: param.DoctorID,
		})
		if errors.Is(err, errors.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		} else if err != nil {
			return nil, errors.Wrap(err, "error when finding doctor")
		}
	}

	patientImages, err := s.patientImageRepository.Find(ctx, params.RepositoryFindPatientImage{
		DoctorID:                  param.DoctorID,
		PatientID:                 param.PatientID,
		RepositoryPaginationParam: param.ToRepositoryPaginationParam(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error when finding patient images")
	}

	res := make([]result.PatientImageBriefInformation, len(patientImages))
	for idx := range res {
		patient, err := s.patientRepository.Find(ctx, params.RepoFindPatient{
			ID: null.NewString(patientImages[idx].PatientID.String(), true),
		})
		if err != nil {
			return nil, errors.Wrap(err, "error when finding patient")
		}

		doctor, err := s.doctorRepository.Find(ctx, params.RepoFindDoctor{
			ID: null.NewString(patientImages[idx].DoctorID.String(), true),
		})
		if err != nil {
			return nil, errors.Wrap(err, "error when finding doctor")
		}

		res[idx].FromPatientModel(patientImages[idx], patient, doctor)
	}

	return res, nil
}
