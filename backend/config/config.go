package config

import (
	"fmt"
	"os"
)

func GetDSN() string {
	// Leemos las variables de entorno (con fallback por si corrés local)
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "tienda")
	dbSSLMode := getEnv("DB_SSLMODE", "false") // en Azure suele ser "require" o "true"

	// DSN (Data Source Name)
	// Si usás Azure MySQL Flexible Server con SSL, agregá ?tls=true o &tls=preferred
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
}

// Helper para leer variables de entorno con default
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
