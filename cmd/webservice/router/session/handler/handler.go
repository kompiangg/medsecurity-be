package handler

import (
	"medsecurity/service/auth"

	"github.com/labstack/echo/v4"
)

type sessionHandler struct {
	authService auth.Service
	e           *echo.Echo
}

func InitSessionHandler(
	e *echo.Echo,
	authService auth.Service,
) sessionHandler {
	return sessionHandler{
		e:           e,
		authService: authService,
	}
}
