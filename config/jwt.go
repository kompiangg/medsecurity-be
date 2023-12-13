package config

type JWT struct {
	DurationInDay int    `yaml:"DurationInDay"`
	Secret        string `yaml:"Secret"`
}
