package main

import (
	"fmt"
	"os"
	"path/filepath"

	"medsecurity/cmd/webservice"
	"medsecurity/config"
	"medsecurity/pkg/db/sqlx"
	_ "medsecurity/pkg/errors"
	"medsecurity/pkg/redis"
	"medsecurity/pkg/validator"
	"medsecurity/repository"
	"medsecurity/service"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer " before the token value
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
		err = nil
	}

	config.UploadFolderPath = tmpDirPath

	db, err := sqlx.InitSQLX(config.Database.URIConnection)
	if err != nil {
		panic(err)
	}

	redis := redis.InitRedis(redis.RedisConfig{
		Hostname: fmt.Sprintf("%s:%s", config.Redis.Hostname, config.Redis.Port),
		Username: config.Redis.Username,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	repository, err := repository.New(
		config,
		db,
		redis,
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
