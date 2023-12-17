package middleware

import (
	"medsecurity/config"
	"medsecurity/pkg/errors"
	"medsecurity/utils/httpx"

	"github.com/labstack/echo/v4"
)

func (m middleware) NotRunInProd() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m.config.Server.Environment == "prod" {
				return httpx.WriteErrorResponse(c, errors.ErrUnauthorized, nil)
			}

			return next(c)
		}
	}
}

func (m middleware) JWTRestricted(jwtType config.JWTType) echo.MiddlewareFunc {
	if jwtType == config.PatientJWT {
		return m.patientJWTMiddleware
	}

	if jwtType == config.DoctorJWT {
		return m.doctorJWTMiddleware
	}

	return m.allRoleMiddleware
}
