package config

import (
	"os"
	"path/filepath"

	"medsecurity/pkg/validator"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database         DatabaseConfig   `yaml:"Database"`
	Server           ServerConfig     `yaml:"Server"`
	Swagger          SwaggerConfig    `yaml:"Swagger"`
	Redis            RedisConfig      `yaml:"Redis"`
	Cloudinary       CloudinaryConfig `yaml:"Cloudinary"`
	JWT              JWTMap           `yaml:"JWT"`
	Permission       Permission       `yaml:"Permission"`
	RSA              RSA              `yaml:"RSA"`
	UploadFolderPath string
}

func InitConfig(validator validator.ValidatorItf) (config Config, err error) {
	fileName, err := filepath.Abs("./etc/config.yaml")
	if err != nil {
		return config, err
	}

	err = loadYamlFile(&config, fileName)
	if err != nil {
		return config, err
	}

	err = validator.Validate(config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func loadYamlFile(cfg *Config, fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	fs, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer fs.Close()

	err = yaml.NewDecoder(fs).Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}
