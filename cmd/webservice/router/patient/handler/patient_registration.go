package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h patientHandler) PatienRegistrationHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServicePatientRegistrationParam
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrBadRequest, "bad request")
		}

		err = h.authService.PatientRegistration(c.Request().Context(), param)
		if errors.Is(err, errors.ErrEmailDuplicated) {
			return httpx.WriteErrorResponse(c, err, "email already registered")
		} else if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when registering patient")
		}

		return httpx.WriteResponse(c, http.StatusCreated, "created")
	}
}
