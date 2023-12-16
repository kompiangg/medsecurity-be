package config

type JWT struct {
	DurationInDay int    `yaml:"DurationInDay"`
	Secret        string `yaml:"Secret"`
}

type JWTType string

const (
	PatientJWT JWTType = "Patient"
	DoctorJWT  JWTType = "Doctor"
)

type JWTMap map[JWTType]JWT
