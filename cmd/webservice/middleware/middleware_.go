package middleware

import (
	"medsecurity/config"
	"medsecurity/pkg/errors"
	"medsecurity/utils/httpx"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (m middleware) NotRunInProd() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m.config.ServerConfig.Environment == "prod" {
				return httpx.WriteErrorResponse(c, errors.ErrUnauthorized, nil, false)
			}

			return next(c)
		}
	}
}

func (m middleware) JWTRestricted(jwtType config.JWTType) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(m.config.JWT[jwtType].Secret),
	})
}
