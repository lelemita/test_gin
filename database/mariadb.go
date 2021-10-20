package database

import (
	"fmt"

	"github.com/lelemita/test_gin/entities"
	"github.com/lelemita/test_gin/utilties"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectInfo struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

var DB *gorm.DB

func init() {
	info := ConnectInfo{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "1234",
		DBName:   "test",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", info.Username, info.Password, info.Host, info.Port, info.DBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utilties.HandleErr(err)
	DB.AutoMigrate(&entities.User{})
}
