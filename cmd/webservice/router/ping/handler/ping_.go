package handler

import (
	"net/http"

	"medsecurity/utils/httpx"

	"github.com/labstack/echo/v4"
)

func (p *pingHandler) PingHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return httpx.WriteResponse(c, http.StatusOK, p.pingService.Ping())
	}
}
