package database

import (
	"fiber-gorm-restapi/config"
	"fiber-gorm-restapi/models"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstamce struct {
	Db *gorm.DB
}

var DB Dbinstamce

func Connect() {
	p := config.Config("DB_PORT") //config function return string, we are parsing str to int here
	port, err := strconv.ParseInt(p, 10, 22)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running AutoMigrations")
	db.AutoMigrate(&models.User{})

	DB = Dbinstamce{
		Db: db,
	}
}
