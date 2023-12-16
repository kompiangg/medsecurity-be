package middleware

import (
	"medsecurity/config"

	"github.com/labstack/echo/v4"
)

type middleware struct {
	config config.Config
}

type Middleware interface {
	NotRunInProd() echo.MiddlewareFunc
	JWTRestricted(jwtType config.JWTType) echo.MiddlewareFunc
}

func New(
	config config.Config,
) middleware {
	return middleware{
		config: config,
	}
}
