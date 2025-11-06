package dao

import (
	"backend/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}

	dial := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})
	gdb, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open error: %v", err)
	}

	cleanup := func() { _ = sqlDB.Close() }
	return gdb, mock, cleanup
}

func TestProductGormRepository_GetByID_OK(t *testing.T) {
	db, mock, cleanup := newMockDB(t)
	defer cleanup()
	repo := NewProductGormRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(1, "Zapa", 100.0)

	// GORM envía: SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?
	q := regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")
	mock.ExpectQuery(q).
		WithArgs(1, 1). // id, limit
		WillReturnRows(rows)

	p, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	assert.Equal(t, uint(1), p.ID)
	assert.Equal(t, "Zapa", p.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductGormRepository_SearchProducts_WithFilter(t *testing.T) {
	db, mock, cleanup := newMockDB(t)
	defer cleanup()
	repo := NewProductGormRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Zapa Roja", 100.0).
		AddRow(2, "Zapa Azul", 120.0)

	q := regexp.QuoteMeta("SELECT * FROM `products` WHERE name LIKE ?")
	mock.ExpectQuery(q).
		WithArgs("%zapa%").
		WillReturnRows(rows)

	list, err := repo.SearchProducts("zapa")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	assert.Len(t, list, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductGormRepository_CreateProduct_OK(t *testing.T) {
	db, mock, cleanup := newMockDB(t)
	defer cleanup()
	repo := NewProductGormRepository(db)

	mock.ExpectBegin()
	// No ates los argumentos: GORM puede incluir timestamps/cols adicionales.
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products`")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := &domain.Product{Name: "Zapa", Price: 100.0}
	err := repo.CreateProduct(p)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
