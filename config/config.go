package config

import (
	"os"
	"path/filepath"

	"medsecurity/pkg/validator"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabaseConfig   DatabaseConfig   `yaml:"Database"`
	ServerConfig     ServerConfig     `yaml:"Server"`
	SwaggerConfig    SwaggerConfig    `yaml:"Swagger"`
	RedisConfig      RedisConfig      `yaml:"Redis"`
	CloudinaryConfig CloudinaryConfig `yaml:"Cloudinary"`
	JWT              JWT              `yaml:"JWT"`
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
