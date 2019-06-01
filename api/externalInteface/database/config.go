package database


type Config struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}
