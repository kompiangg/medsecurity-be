package handler

import (
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

type otpHandler struct {
	e                   *echo.Echo
	patientImageService patient_image.Service
}

func InitOTPHandler(
	e *echo.Echo,
	patientImageService patient_image.Service,
) otpHandler {
	return otpHandler{
		e:                   e,
		patientImageService: patientImageService,
	}
}
