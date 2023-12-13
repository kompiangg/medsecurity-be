package config

type DatabaseConfig struct {
	LongTermStorageDSN   string   `yaml:"LongTermStorageDSN"`
	TimeSeriesStorageDSN []string `yaml:"TimeSeriesStorageDSN"`
	TimeSeriesKey        []string `yaml:"TimeSeriesKey"`
}
