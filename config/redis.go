package config

type RedisConfig struct {
	Hostname string `yaml:"Hostname"`
	Username string `yaml:"Username"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"Db"`
}
