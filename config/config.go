package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"urlshortener/domain"
)

var DB *gorm.DB

func InitDB() {
	ENV := LoadENV()
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV["username"], ENV["password"], ENV["host"], ENV["database"])
	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		log.Fatal("Failed connect to database : ", e.Error())
	}

	//// migration
	InitMigration()
}

func InitMigration() {
	err := DB.AutoMigrate(&domain.Url{}, &domain.User{})
	if err != nil {
		log.Fatal("Error migration : ", err.Error())
	}
}

func LoadENV() map[string]string {
	// load file env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error when load env : ", err.Error())
	} else {
		log.Println("Success load env")
	}
	// mapping read env
	return map[string]string{
		"host":      os.Getenv("DB_HOST"),
		"username":  os.Getenv("DB_USERNAME"),
		"password":  os.Getenv("DB_PASSWORD"),
		"database":  os.Getenv("DB_DATABASE"),
		"jwtSecret": os.Getenv("JWT_SECRET"),
	}

}
