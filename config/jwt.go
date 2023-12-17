package config

type JWT struct {
	DurationInDay int    `yaml:"DurationInDay"`
	Secret        string `yaml:"Secret"`
}

type JWTType string

const (
	DoctorJWT  JWTType = "Doctor"
	PatientJWT JWTType = "Patient"
	AllRoleJWT JWTType = "AllRole"
)

type JWTMap map[JWTType]JWT
