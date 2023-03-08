package configService

import (
	logService "backend-skeleton-golang/commons/app/services/log-service"
	configInfra "backend-skeleton-golang/commons/infra/godotenv"
	"strconv"
)

func GetJwtSecret() string {
	return configInfra.Get("JWT_SECRET")
}

func GetJwtExt() int {
	value, err := strconv.Atoi(configInfra.Get("JWT_EXP"))
	if err != nil {
		logService.Error(err.Error())
		return 14440
	}
	return value
}

func GetDbConfig() string {
	return configInfra.Get("DB_CONFIG")
}

func GetSmtpHost() string {
	return configInfra.Get("SMTP_HOST")
}

func GetSmtpPort() int {
	value, err := strconv.Atoi(configInfra.Get("SMTP_PORT"))
	if err != nil {
		panic("SMTP_PORT must be a int")
	}
	return value
}

func GetSmtpUser() string {
	return configInfra.Get("SMTP_USER")
}

func GetSmtpPass() string {
	return configInfra.Get("SMTP_PASS")
}

func GetSmtpFrom() string {
	return configInfra.Get("SMTP_FROM")
}
