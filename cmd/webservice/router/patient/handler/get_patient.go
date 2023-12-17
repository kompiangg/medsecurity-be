package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/constant"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientHandler) GetPatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceFindPatient
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "bad request")
		}

		role, err := httpx.GetRoleAccountFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error at parsing role")
		}

		if role == constant.PatientRole {
			jwtClaim, err := httpx.GetJWTClaimsFromContext(c)
			if err != nil {
				return httpx.WriteErrorResponse(c, err, "error at parsing jwt claim")
			}

			patientID, ok := jwtClaim["sub"].(string)
			if !ok {
				return httpx.WriteErrorResponse(c, err, "error at parsing patient id")
			}

			if patientID != param.PatientID {
				return httpx.WriteErrorResponse(c, errors.ErrUnauthorized, "you are not authorized to access this resource")
			}
		}

		patient, err := h.patientService.FindPatientByID(c.Request().Context(), param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error at GetPatient")
		}

		return httpx.WriteResponse(c, http.StatusOK, patient)
	}
}
