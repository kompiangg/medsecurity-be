package middleware

import (
	"medsecurity/config"
	"medsecurity/type/constant"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	config               config.Config
	patientJWTMiddleware echo.MiddlewareFunc
	doctorJWTMiddleware  echo.MiddlewareFunc
	allRoleMiddleware    echo.MiddlewareFunc
}

type Middleware interface {
	NotRunInProd() echo.MiddlewareFunc
	JWTRestricted(jwtType config.JWTType) echo.MiddlewareFunc
}

func New(
	configParam config.Config,
) middleware {
	patientJWTMiddlewareConfig := echojwt.Config{
		SigningKey: []byte(configParam.JWT[config.PatientJWT].Secret),
		SuccessHandler: func(c echo.Context) {
			c.Set(constant.ContextKeyRole, constant.PatientRole)
		},
	}

	doctorJWTMiddlewareConfig := echojwt.Config{
		SigningKey: []byte(configParam.JWT[config.DoctorJWT].Secret),
		SuccessHandler: func(c echo.Context) {
			c.Set(constant.ContextKeyRole, constant.DoctorRole)
		},
	}

	allRoleMiddlewareConfig := func(next echo.HandlerFunc) echo.HandlerFunc {
		doctorMiddleware := doctorJWTMiddlewareConfig
		doctorMiddleware.ErrorHandler = func(c echo.Context, err error) error {
			return echojwt.WithConfig(patientJWTMiddlewareConfig)(next)(c)
		}

		return echojwt.WithConfig(doctorMiddleware)(next)
	}

	return middleware{
		config:               configParam,
		patientJWTMiddleware: echojwt.WithConfig(patientJWTMiddlewareConfig),
		doctorJWTMiddleware:  echojwt.WithConfig(doctorJWTMiddlewareConfig),
		allRoleMiddleware:    allRoleMiddlewareConfig,
	}
}
