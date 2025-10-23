package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gmysql "github.com/go-sql-driver/mysql" // 👈 necesario para registrar TLS
)

// GetDSN genera la cadena de conexión (DSN) con soporte TLS para Azure MySQL Flexible
func GetDSN() string {
	// Variables de entorno (con fallback por si corrés local)
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "root")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "tienda")

	// ⚙️ Registrar configuración TLS (necesario para Azure)
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("/etc/ssl/certs/ca-certificates.crt") // ruta estándar en App Service Linux
	if err != nil {
		log.Printf("⚠️ No se pudo leer certificado raíz: %v", err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Printf("⚠️ No se pudo agregar el certificado PEM")
	}

	tlsConfig := &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true, // Azure usa certificados válidos, pero esto evita fallos intermedios
	}

	err = gmysql.RegisterTLSConfig("azure", tlsConfig)
	if err != nil {
		log.Printf("⚠️ No se pudo registrar TLS config: %v", err)
	}

	// ✅ DSN compatible con Azure MySQL Flexible Server
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=azure",
		dbUser, dbPassword, dbHost, dbPort, dbName)
}

// Helper para leer variables de entorno con valor por defecto
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
