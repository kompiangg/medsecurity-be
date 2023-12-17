package otp

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/otp/handler"
	"medsecurity/config"
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	patientImageService patient_image.Service,
	middleware middleware.Middleware,
) {
	handler := handler.InitOTPHandler(e, patientImageService)

	e.POST(V1PatientRequestGetImage, handler.PatientRequestGetImage, middleware.JWTRestricted(config.AllRoleJWT))
}
