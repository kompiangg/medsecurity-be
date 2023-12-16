package upload

import (
	"medsecurity/cmd/webservice/router/ping/handler"
	"medsecurity/config"
	"medsecurity/service/ping"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	fileService ping.ServiceItf,
	config config.Config,
) {
	pingHandler := handler.InitPingHandler(e, fileService)

	e.GET(V1PingPath, pingHandler.PingHandler())
}
