package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientHandler) GetPatientsImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceFindPatientImage
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrBadRequest, "bad request", false)
		}

		claims, err := httpx.GetJWTClaimsFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil, false)
		}

		accountID, ok := claims["sub"].(string)
		if !ok {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil, false)
		}

		role, err := httpx.GetRoleAccountFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil, false)
		}

		param.Role = role
		param.AccountID = accountID

		res, err := h.patientService.FindPatientImageBriefInformation(c.Request().Context(), param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, nil, false)
		}

		return httpx.WriteResponse(c, http.StatusOK, res)
	}
}
