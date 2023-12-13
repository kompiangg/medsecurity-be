package main

import (
	"fmt"
	"os"
	"path/filepath"

	"medsecurity/cmd/webservice"
	"medsecurity/config"
	"medsecurity/pkg/db/sqlx"
	_ "medsecurity/pkg/errors"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer " before the token value
func main() {
	validator, err := validator.New()
	if err != nil {
		panic(err)
	}

	config, err := config.InitConfig(&validator)
	if err != nil {
		panic(err)
	}

	tmpDirPath, err := filepath.Abs("./tmp")
	if err != nil {
		panic(err)
	}

	if err := os.Mkdir(tmpDirPath, 0755); os.IsExist(err) {
		fmt.Println("The directory named", tmpDirPath, "exists")
	} else {
		fmt.Println("The directory named", tmpDirPath, "does not exist")
	}

	config.UploadFolderPath = tmpDirPath

	longTermStorage, err := sqlx.InitSQLX(config.DatabaseConfig.LongTermStorageDSN)
	if err != nil {
		panic(err)
	}

	// timeseriesdatabase, err := timeseriesdatabase.InitTimeSeriesDatabase(config.DatabaseConfig)
	// if err != nil {
	// 	panic(err)
	// }

	// redis := redis.InitRedis(redis.RedisConfig{
	// 	Hostname: fmt.Sprintf("%s:%s", config.RedisConfig.Hostname, config.RedisConfig.Port),
	// 	Username: config.RedisConfig.Username,
	// 	Password: config.RedisConfig.Password,
	// 	DB:       config.RedisConfig.DB,
	// })

	// cloudinary, err := cloudinary.InitCloudinary(cloudinary.CloudinaryConfig{
	// 	APIKey:    config.CloudinaryConfig.APIKey,
	// 	APISecret: config.CloudinaryConfig.APISecret,
	// 	CloudName: config.CloudinaryConfig.CloudName,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	repository, err := repository.New(
		config,
		longTermStorage,
		// redis,
		// cloudinary,
	)
	if err != nil {
		panic(err)
	}

	service, err := service.New(
		repository,
		config,
		&validator,
	)
	if err != nil {
		panic(err)
	}

	err = webservice.InitWebService(
		service,
		config,
		&validator,
	)
	if err != nil {
		panic(err)
	}
}
