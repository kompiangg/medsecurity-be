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
		Port:                 config.Server.Port,
		Environment:          config.Server.Environment,
		WhiteListAllowOrigin: config.Server.WhiteListAllowOrigin,
	})
	if err != nil {
		return err
	}

	echo.Echo.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: config.Server.WhiteListAllowOrigin,
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
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API", config.Swagger.Title)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.Swagger.Hostname, config.Swagger.Port)
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Version = config.Swagger.Version
	docs.SwaggerInfo.Description = config.Swagger.Description
	docs.SwaggerInfo.Schemes = config.Swagger.Schemes
}
