package handler

import (
	"medsecurity/service/auth"
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

type patientImageHandler struct {
	authService         auth.Service
	patientImageService patient_image.Service
	e                   *echo.Echo
}

func InitImagePatientHandler(
	e *echo.Echo,
	authService auth.Service,
	patientImageService patient_image.Service,
) patientImageHandler {
	return patientImageHandler{
		e:                   e,
		authService:         authService,
		patientImageService: patientImageService,
	}
}
