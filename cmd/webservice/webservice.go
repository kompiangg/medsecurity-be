package webservice

import (
	"fmt"

	inmiddleware "medsecurity/cmd/webservice/middleware"
	"medsecurity/cmd/webservice/router"
	"medsecurity/config"
	"medsecurity/docs"
	"medsecurity/pkg/http"
	"medsecurity/pkg/validator"
	"medsecurity/service"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitWebService(
	service service.Service,
	config config.Config,
	validator validator.ValidatorItf,
) error {
	echo, err := http.InitEchoServer(&http.ServerConfig{
		Port:                 config.ServerConfig.Port,
		Environment:          config.ServerConfig.Environment,
		WhiteListAllowOrigin: config.ServerConfig.WhiteListAllowOrigin,
	})
	if err != nil {
		return err
	}

	echo.Echo.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: config.ServerConfig.WhiteListAllowOrigin,
		},
	))

	middleware := inmiddleware.New(config)

	// Swagger
	initSwagger(config)
	echo.Echo.GET("/swagger/*", echoSwagger.WrapHandler, middleware.NotRunInProd())

	// Register All Handler
	router.RegisterHandler(echo.Echo, service, config, middleware)

	// Start HTTP Server
	err = echo.ServeHTTP()
	if err != nil {
		return err
	}

	return nil
}

func initSwagger(config config.Config) {
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API", config.SwaggerConfig.Title)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.SwaggerConfig.Hostname, config.SwaggerConfig.Port)
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Version = config.SwaggerConfig.Version
	docs.SwaggerInfo.Description = config.SwaggerConfig.Description
	docs.SwaggerInfo.Schemes = config.SwaggerConfig.Schemes
}
