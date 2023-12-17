package handler

import (
	"medsecurity/pkg/errors"
	"medsecurity/type/params"
	"medsecurity/utils/httpx"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h sessionHandler) DoctorLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param params.ServiceDoctorLoginParam
		err := c.Bind(&param)
		if err != nil {
			return httpx.WriteErrorResponse(c, errors.ErrBadRequest, "bad request")
		}

		token, err := h.authService.DoctorLogin(c.Request().Context(), param)
		if errors.Is(err, errors.ErrAccountNotFound) {
			return httpx.WriteErrorResponse(c, err, "account not found")
		} else if err != nil {
			return httpx.WriteErrorResponse(c, err, "error when login")
		}

		return httpx.WriteResponse(c, http.StatusOK, token)
	}
}
