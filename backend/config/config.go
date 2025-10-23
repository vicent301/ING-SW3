package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	gmysql "github.com/go-sql-driver/mysql" // 🔹 alias para registrar TLS
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUser := os.Getenv("DBUser")
	dbPass := os.Getenv("DBPassword")
	dbHost := os.Getenv("DBHost")
	dbPort := os.Getenv("DBPort")
	dbName := os.Getenv("DBName")

	// ⚙️ 1. Registrar configuración TLS requerida por Azure
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("/etc/ssl/certs/ca-certificates.crt") // ruta estándar en Linux App Service
	if err != nil {
		log.Fatalf("❌ No se pudo leer el certificado raíz: %v", err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatalf("❌ No se pudo agregar el certificado PEM")
	}

	tlsConfig := &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
	}

	// 🔹 Registrar TLS en el driver real
	err = gmysql.RegisterTLSConfig("azure", tlsConfig)
	if err != nil {
		log.Fatalf("❌ No se pudo registrar TLS config: %v", err)
	}

	// ⚙️ 2. Construir DSN compatible con Azure
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=azure",
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Printf("Intentando conectar a: %s", dbHost)

	// ⚙️ 3. Probar conexión simple
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error al abrir conexión: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("❌ Error al conectar a la base de datos: %v", err)
	}

	// ⚙️ 4. Iniciar GORM con el mismo DSN
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error al inicializar GORM: %v", err)
	}

	log.Println("✅ Conectado exitosamente a la base de datos con TLS (Azure MySQL Flexible)")
}
