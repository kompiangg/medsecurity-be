package service

import (
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service/auth"
	"medsecurity/service/ping"
)

type Service struct {
	Ping ping.ServiceItf
	Auth auth.Service
}

func New(
	repository repository.Repository,
	config config.Config,
	validator validator.ValidatorItf,
) (Service, error) {
	pingService := ping.New()
	authService := auth.New(repository.Auth, repository.Doctor, repository.Patient, validator)

	return Service{
		Ping: pingService,
		Auth: authService,
	}, nil
}
