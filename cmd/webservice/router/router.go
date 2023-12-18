package router

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/access_request"
	"medsecurity/cmd/webservice/router/doctor"
	"medsecurity/cmd/webservice/router/otp"
	"medsecurity/cmd/webservice/router/patient"
	"medsecurity/cmd/webservice/router/patient_image"
	ping "medsecurity/cmd/webservice/router/ping"
	"medsecurity/cmd/webservice/router/session"
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
	patient.InitHandler(echo, service.Auth, service.Patient, config, middleware)
	patient_image.InitHandler(echo, service.Auth, service.Patient, service.PatientImage, config, middleware)
	doctor.InitHandler(echo, service.Auth, config, middleware)
	session.InitHandler(echo, service.Auth, config, middleware)
	otp.InitHandler(echo, service.PatientImage, middleware)
	access_request.InitHandler(echo, service.PatientImage, middleware)
}
