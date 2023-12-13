package config

type ServerConfig struct {
	Port                 string   `yaml:"Port" validate:"required,notblank,numeric"`
	Environment          string   `yaml:"Environment" validate:"required,notblank,oneof=dev prod"`
	WhiteListAllowOrigin []string `yaml:"WhiteListCORS" validate:"required,notblank"`
}
