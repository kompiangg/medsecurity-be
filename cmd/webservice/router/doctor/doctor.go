package doctor

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/doctor/handler"
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
	handler := handler.InitAuthHandler(e, authService)

	e.POST(V1DoctorRegistrationPath, handler.DoctorRegistrationHandler())
}
