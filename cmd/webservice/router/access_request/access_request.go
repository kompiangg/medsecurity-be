package access_request

import (
	"medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router/access_request/handler"
	"medsecurity/config"
	"medsecurity/service/patient_image"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	patientImageService patient_image.Service,
	middleware middleware.Middleware,
) {
	handler := handler.InitAccessRequestHandler(e, patientImageService)

	e.POST("/v1/access-request", handler.RequestAccess(), middleware.JWTRestricted(config.PatientJWT))
}
