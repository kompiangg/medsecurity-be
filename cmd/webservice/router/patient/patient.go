package patient

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/patient/handler"
	"medsecurity/config"
	"medsecurity/service/auth"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	authService auth.Service,
	config config.Config,
	middleware middleware.Middleware,
) {
	handler := handler.InitPatientHandler(e, authService)

	e.POST(V1PatientRegistrationPath, handler.PatienRegistrationHandler())
}
