package handler

import (
	"medsecurity/service/auth"
	"medsecurity/service/patient"

	"github.com/labstack/echo/v4"
)

type patientHandler struct {
	authService    auth.Service
	patientService patient.Service
	e              *echo.Echo
}

func InitPatientHandler(
	e *echo.Echo,
	authService auth.Service,
	patientService patient.Service,
) patientHandler {
	return patientHandler{
		e:              e,
		authService:    authService,
		patientService: patientService,
	}
}
