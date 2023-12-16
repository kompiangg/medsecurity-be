package handler

import (
	"medsecurity/service/auth"

	"github.com/labstack/echo/v4"
)

type patientHandler struct {
	authService auth.Service
	e           *echo.Echo
}

func InitPatientHandler(
	e *echo.Echo,
	authService auth.Service,
) patientHandler {
	return patientHandler{
		e:           e,
		authService: authService,
	}
}
