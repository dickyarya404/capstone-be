package database

import (
	"fmt"
	"log"

	"github.com/sawalreverr/recything/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	DB *gorm.DB
}

var (
	dbInstance *mysqlDatabase
)

func NewMySQLDatabase(conf *config.Config) Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	log.Println("Successfully connected to database:", conf.DB.DBName)

	dbInstance = &mysqlDatabase{DB: db}
	return dbInstance
}

func (m *mysqlDatabase) GetDB() *gorm.DB {
	return dbInstance.DB
}
