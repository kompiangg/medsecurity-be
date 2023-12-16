package session

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/session/handler"
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
	handler := handler.InitSessionHandler(e, authService)

	e.POST(V1DoctorLogin, handler.DoctorLogin())
	e.POST(V1PatientLogin, handler.PatientLogin())
}
