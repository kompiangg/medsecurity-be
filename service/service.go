package service

import (
	"medsecurity/config"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service/ping"
)

type Service struct {
	Ping ping.ServiceItf
}

func New(
	repository repository.Repository,
	config config.Config,
	validator validator.ValidatorItf,
) (Service, error) {
	pingService := ping.New()

	return Service{
		Ping: pingService,
	}, nil
}
