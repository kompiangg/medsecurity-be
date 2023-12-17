package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/constant"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientHandler) GetAllPatients() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceFindAllPatients
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "bad request")
		}

		role, err := httpx.GetRoleAccountFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error at parsing role")
		}

		if role != constant.DoctorRole {
			return httpx.WriteErrorResponse(c, errors.ErrUnauthorized, "you are not authorized to access this resource")
		}

		patients, err := h.patientService.FindPatients(c.Request().Context(), param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error at GetAllPatients")
		}

		return httpx.WriteResponse(c, http.StatusOK, patients)
	}
}
