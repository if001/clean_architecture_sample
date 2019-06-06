package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBConnection struct {
	DB *gorm.DB
}

func NewSqlConnection() DBConnection {
	config, err := LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	dbconf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.DB)
	fmt.Println("hogehgoe:", dbconf)
	db, err := gorm.Open("mysql", dbconf)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	// defer db.Close()
	return DBConnection{DB: db}
}
