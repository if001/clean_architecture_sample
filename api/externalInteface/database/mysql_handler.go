package database
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/BurntSushi/toml"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewSqlHandler() *gorm.DB {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	dbconf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DB)
	db, err := gorm.Open("mysql", dbconf)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	// defer db.Close()
	return db
}

