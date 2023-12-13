package config

type Permission struct {
	Admin string `yaml:"Admin"`
	Owner string `yaml:"Owner"`
	User  string `yaml:"User"`
}
