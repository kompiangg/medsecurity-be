package handler

import (
	"medsecurity/service/ping"

	"github.com/labstack/echo/v4"
)

type pingHandler struct {
	pingService ping.ServiceItf
	e           *echo.Echo
}

func InitPingHandler(
	e *echo.Echo,
	pingService ping.ServiceItf,
) pingHandler {
	return pingHandler{
		e:           e,
		pingService: pingService,
	}
}
