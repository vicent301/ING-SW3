package dao_test

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"testing"
)

func TestProductDAO_CreateAndSearch(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// create product vía DAO plano si lo tenés, si no seed directo
	if err := dao.CreateProduct(domain.Product{
		Name: "Nike Air", Price: 100, Stock: 7, Description: "Zapa", ImageURL: "",
	}); err != nil {
		t.Fatalf("CreateProduct: %v", err)
	}
	if err := database.DB.Create(&dao.Product{
		Name: "Adidas Samba", Price: 90, Stock: 5,
	}).Error; err != nil {
		t.Fatalf("seed: %v", err)
	}

	// search vacío => trae todo
	all, err := dao.SearchProducts("")
	if err != nil {
		t.Fatalf("SearchProducts: %v", err)
	}
	if len(all) < 2 {
		t.Fatalf("esperaba al menos 2 productos, got %d", len(all))
	}

	// search por “Nike”
	nike, err := dao.SearchProducts("Nike")
	if err != nil {
		t.Fatalf("SearchProducts Nike: %v", err)
	}
	if len(nike) == 0 {
		t.Fatalf("esperaba match para 'Nike'")
	}
}

func TestProductDAO_GetByID_NotFound(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	if _, err := dao.GetProductByID(99999); err == nil {
		t.Fatalf("esperaba error por producto inexistente")
	}
}

func TestProductDAO_SearchProducts_NoResults(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	// sin sembrar nada => buscar devuelve vacío
	res, err := dao.SearchProducts("no-deberia-existir")
	if err != nil {
		t.Fatalf("SearchProducts: %v", err)
	}
	if len(res) != 0 {
		t.Fatalf("esperaba 0 resultados, got %d", len(res))
	}
}
