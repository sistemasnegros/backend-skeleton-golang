package godotenvInfra

import (
	logService "backend-skeleton-golang/commons/app/services/log-service"
	configDomain "backend-skeleton-golang/commons/domain/config"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load("../.env")
	if err != nil {
		logService.Error("error loading .env")
	}

	// Validate field
	config := &configDomain.Config{}
	val := reflect.ValueOf(config).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i).Name
		if Get(field) == "" {
			logService.Error("undefine in env: " + field)
			os.Exit(1)
		}
		logService.Info("load env: " + field)
	}

}

func Get(key string) string {
	value := os.Getenv(key)
	return value
}
