package patient

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/patient/handler"
	"medsecurity/config"
	"medsecurity/service/auth"
	"medsecurity/service/patient"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	authService auth.Service,
	patientService patient.Service,
	paramConfig config.Config,
	middleware middleware.Middleware,
) {
	handler := handler.InitPatientHandler(e, authService, patientService)

	e.GET("/v1/patient/:patient_id", handler.GetPatient(), middleware.JWTRestricted(config.AllRoleJWT))
	e.GET("/v1/patient", handler.GetAllPatients(), middleware.JWTRestricted(config.AllRoleJWT))
	e.POST("/v1/patient", handler.PatienRegistrationHandler())
}
