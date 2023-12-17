package config

type AES struct {
	Salt   string `yaml:"Salt"`
	Secret string `yaml:"Secret"`
}
