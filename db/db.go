package db

import (
	"log"

	"github.com/coala/corobo-ng/config"
	"github.com/coala/corobo-ng/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

func Init() {
	var db = DB

	config := config.GetConfig()

	host := config.Database.Host
	port := config.Database.Port
	name := config.Database.Name
	user := config.Database.User
	password := config.Database.Password
	driver := config.Database.Driver

	if driver == "postgres" {
		dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + name + " sslmode=disable password=" + password
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			DBErr = err
			log.Fatalf("Database error: %v", err)
		}
	} else {
		log.Fatalf("Database driver %s is not supported", driver)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func GetDBErr() error {
	return DBErr
}
