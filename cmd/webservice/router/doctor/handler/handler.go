package handler

import (
	"medsecurity/service/auth"

	"github.com/labstack/echo/v4"
)

type doctorHandler struct {
	authService auth.Service
	e           *echo.Echo
}

func InitAuthHandler(
	e *echo.Echo,
	authService auth.Service,
) doctorHandler {
	return doctorHandler{
		e:           e,
		authService: authService,
	}
}
