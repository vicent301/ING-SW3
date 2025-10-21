package config

import "fmt"

var (
	DBUser     = "root"
	DBPassword = "root"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "tienda"
)

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DBUser, DBPassword, DBHost, DBPort, DBName)
}
