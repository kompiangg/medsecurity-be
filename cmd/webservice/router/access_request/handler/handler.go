package handler

import (
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

type accessRequestHandler struct {
	e                   *echo.Echo
	patientImageService patient_image.Service
}

func InitAccessRequestHandler(
	e *echo.Echo,
	patientImageService patient_image.Service,
) accessRequestHandler {
	return accessRequestHandler{
		e:                   e,
		patientImageService: patientImageService,
	}
}
