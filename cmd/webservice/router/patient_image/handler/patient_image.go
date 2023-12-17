package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientImageHandler) GetPatientsImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceFindPatientImage
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrBadRequest, "bad request")
		}

		claims, err := httpx.GetJWTClaimsFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		accountID, ok := claims["sub"].(string)
		if !ok {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		role, err := httpx.GetRoleAccountFromContext(c)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrInternalServer, nil)
		}

		param.Role = role
		param.AccountID = accountID

		res, err := h.patientImageService.FindBriefInformation(c.Request().Context(), param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, nil)
		}

		return httpx.WriteResponse(c, http.StatusOK, res)
	}
}
