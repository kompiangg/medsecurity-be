package params

import (
	"github.com/volatiletech/null/v9"
)

type RepositoryFindPatientImage struct {
	PatientID null.String
	DoctorID  null.String

	RepositoryPaginationParam
}

type ServiceFindPatientImage struct {
	PatientID null.String `query:"patient_id"`
	DoctorID  null.String `query:"doctor_id"`

	Role      string `param:"-"`
	AccountID string `param:"-"`

	ServicePaginationParam
}
