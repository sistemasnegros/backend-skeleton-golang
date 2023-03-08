package gormInfra

import (
	configService "backend-skeleton-golang/commons/app/services/config-service"
	usersModel "backend-skeleton-golang/users/infra/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func Connect() *gorm.DB {
	Db, err := gorm.Open(sqlite.Open("../"+configService.GetDbConfig()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&usersModel.User{})

	return Db
}


func New() *gorm.DB{
	return Connect()
}
