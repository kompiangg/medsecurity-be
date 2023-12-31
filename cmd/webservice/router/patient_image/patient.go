package patient_image

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/patient_image/handler"
	"medsecurity/config"
	"medsecurity/service/auth"
	"medsecurity/service/patient"
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	authService auth.Service,
	patientService patient.Service,
	patientImageService patient_image.Service,
	paramConfig config.Config,
	middleware middleware.Middleware,
) {
	handler := handler.InitImagePatientHandler(
		e,
		authService,
		patientImageService,
	)

	e.GET("/v1/patient-image/doctor/:permission_id", handler.DoctorGetDecryptedImage(), middleware.JWTRestricted(config.AllRoleJWT))
	e.GET(V1FindDecryptedPatientImage, handler.GetDecryptedImage(), middleware.JWTRestricted(config.AllRoleJWT))
	e.POST(V1InsertPatientsImage, handler.Insert(), middleware.JWTRestricted(config.AllRoleJWT))
	e.GET(V1FindPatientsImage, handler.GetPatientsImage(), middleware.JWTRestricted(config.AllRoleJWT))
}
