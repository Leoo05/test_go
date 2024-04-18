package postgreDatabase

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func GetDB() *gorm.DB {
	if instance == nil {
		fmt.Println("creating instance of repository")
		instance, _ = InitDB()
	}
	return instance
}

func InitDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	dsn := connString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to dabase: %s", err.Error())
	}
	return db, nil
}
