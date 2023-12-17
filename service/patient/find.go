package patient

import (
	"context"
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/type/result"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v9"
)

func (s service) FindPatientByID(ctx context.Context, param params.ServiceFindPatient) (result.ServiceGetDetailPatient, error) {
	err := s.validator.Validate(param)
	if err != nil {
		return result.ServiceGetDetailPatient{}, err
	}

	patientID, err := uuid.Parse(param.PatientID)
	if err != nil {
		return result.ServiceGetDetailPatient{}, errors.Wrap(err, "error at FindPatientByID")
	}

	patient, err := s.patientRepository.Find(ctx, params.RepoFindPatient{
		ID: null.NewString(patientID.String(), true),
	})
	if errors.Is(err, errors.ErrRecordNotFound) {
		return result.ServiceGetDetailPatient{}, errors.ErrAccountNotFound
	} else if err != nil {
		return result.ServiceGetDetailPatient{}, errors.Wrap(err, "error at FindPatientByID")
	}

	res := result.ServiceGetDetailPatient{}
	res.FromPatientModel(patient)

	return res, nil
}

func (s service) FindPatients(ctx context.Context, param params.ServiceFindAllPatients) ([]result.ServiceGetAllPatients, error) {
	patients, err := s.patientRepository.FindAll(ctx, params.RepoFindAllPatients{
		RepositoryPagination: param.CreatePagination(),
	})
	if err != nil {
		return []result.ServiceGetAllPatients{}, errors.Wrap(err, "error at FindPatients")
	}

	res := make([]result.ServiceGetAllPatients, len(patients))
	for i := range patients {
		res[i].FromPatientModel(patients[i])
	}

	return res, nil
}
