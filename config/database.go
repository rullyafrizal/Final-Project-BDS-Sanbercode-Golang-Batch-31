package config

import (
	"fmt"
	"log"
	"os"

	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func BuildDBConfig() *DbConfig {
	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}

func DbUrl(dbConfig *DbConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
}

func ConnectPostgres() *gorm.DB {
	var dsn string = DbUrl(BuildDBConfig())

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	db.SetupJoinTable(&models.Post{}, "Votes", &models.Vote{})

	db.AutoMigrate(
		&models.Tag{},
		&models.Role{},
		&models.User{},
		&models.Post{},
		&models.PostImage{},
		&models.Review{},
	)

	return db
}
