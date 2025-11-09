package database_test

import (
	"backend/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
)

func TestConnectWithDialector_UsesInjectedConnAndSetsGlobalDB(t *testing.T) {
	// Creamos un *sql.DB falso
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer sqlDB.Close()
	_ = mock // si querés, podés setear expectativas de ping/query aquí

	// Dialector de MySQL usando la conexión mockeada
	dial := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // evita consulta de versión
	})

	// Llamamos a la función de test
	db := database.ConnectWithDialector(dial)
	if db == nil {
		t.Fatalf("ConnectWithDialector devolvió nil")
	}

	// Verificamos que la global quede seteada
	if database.DB == nil {
		t.Fatalf("database.DB no fue seteada")
	}

	// Abrimos la conexión subyacente y comprobamos que sea la misma
	gotSQL, err := database.DB.DB()
	if err != nil {
		t.Fatalf("DB.DB(): %v", err)
	}
	if gotSQL != sqlDB {
		t.Fatalf("la conexión subyacente no coincide con la de sqlmock")
	}
}
