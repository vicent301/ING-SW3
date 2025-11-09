package config

import (
	"strings"
	"testing"
)

func TestGetDSN_BuildsFromEnvAndHasTLS(t *testing.T) {
	t.Setenv("DB_USER", "u")
	t.Setenv("DB_PASSWORD", "p")
	t.Setenv("DB_HOST", "host.example.com")
	t.Setenv("DB_PORT", "3307")
	t.Setenv("DB_NAME", "dbname")

	dsn := GetDSN()

	// Contenido esperado básico
	wantParts := []string{
		"u:p@tcp(host.example.com:3307)/dbname",
		"parseTime=true",
		"tls=", // que use la config TLS registrada
	}
	for _, w := range wantParts {
		if !strings.Contains(dsn, w) {
			t.Fatalf("dsn no contiene %q; got=%s", w, dsn)
		}
	}
}
