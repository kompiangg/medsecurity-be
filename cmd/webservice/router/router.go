package router

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/doctor"
	"medsecurity/cmd/webservice/router/patient"
	ping "medsecurity/cmd/webservice/router/ping"
	"medsecurity/config"
	"medsecurity/service"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(
	echo *echo.Echo,
	service service.Service,
	config config.Config,
	middleware middleware.Middleware,
) {
	ping.InitHandler(echo, service.Ping, config)
	patient.InitHandler(echo, service.Auth, config, middleware)
	doctor.InitHandler(echo, service.Auth, config, middleware)
}
