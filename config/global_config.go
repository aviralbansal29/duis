package config

import (
	"sync"

	"github.com/aviralbansal29/duis/constant"
	"github.com/aviralbansal29/duis/log"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// GlobalInstance should be reused.
type GlobalInstance struct {
	env        *viper.Viper
	database   *gorm.DB
	validation *validator.Validate
	resty      *resty.Client
}

var globalInstance GlobalInstance
var globalInstanceOnce sync.Once

// InitiateGlobalInstance initiates the config data
func InitiateGlobalInstance() {
	globalInstanceOnce.Do(func() {
		env := loadConfig(constant.EnvVariablePath, constant.EnvFileName, constant.EnvFileExtension)
		database := loadDBHandler(env)
		log.SetupLogger()
		globalInstance = GlobalInstance{
			env:        env,
			database:   database,
			validation: validator.New(),
			resty:      resty.New(),
		}
	})
}

// GetEnv returns database instance
func GetEnv() *viper.Viper {
	return globalInstance.env
}

// DatabaseHandler returns postgres main database handler
func DatabaseHandler() *gorm.DB {
	return globalInstance.database
}

// Validator returns validator object
func Validator() *validator.Validate {
	return globalInstance.validation
}

// RestyClient returns client for resty library
func RestyClient() *resty.Client {
	return globalInstance.resty
}
