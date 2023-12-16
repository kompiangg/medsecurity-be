package config

type RSA struct {
	KeySize int `yaml:"KeySize" validate:"required"`
}
