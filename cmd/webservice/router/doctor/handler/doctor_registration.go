package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h doctorHandler) DoctorRegistrationHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceDoctorRegistrationParam
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when binding param", false)
		}

		err = h.authService.DoctorRegistration(c.Request().Context(), param)
		if errors.Is(err, errors.ErrEmailDuplicated) {
			return httpx.WriteErrorResponse(c, err, "email already registered", false)
		} else if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when registering doctor", true)
		}

		return httpx.WriteResponse(c, http.StatusCreated, "created")
	}
}
