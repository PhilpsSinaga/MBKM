package configuration

import (
	"GO_MINIPROJECT/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() (*gorm.DB, error) {
	config := Config{
		DB_Username: "root",
		DB_Password: "password",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "gym",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB, err
}

func GetDB() *gorm.DB {
	return DB
}

func InitialMigration(DB *gorm.DB) error {
	return DB.AutoMigrate(&model.User{}, &model.Admins{}, &model.Class{}, &model.Membership{}, &model.Payment_method{}, &model.Trainer{})
}
