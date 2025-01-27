package db

import (
	"fmt"
	"photo-sharing-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(
	postgresUser string,
	postgresHost string,
	postgresPassword string,
	postgresDatabase string,
	postgresPort string,
) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC",
		postgresHost,
		postgresUser,
		postgresPassword,
		postgresDatabase,
		postgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Users{})

	return db
}
