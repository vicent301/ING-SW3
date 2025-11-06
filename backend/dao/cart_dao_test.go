package dao_test

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"testing"
)

func seedUser(t *testing.T, email string) *domain.User {
	t.Helper()
	u := domain.User{Name: "Vic", Email: email, Password: "x"}
	if err := dao.CreateUser(u); err != nil {
		t.Fatalf("seed user: %v", err)
	}
	got, err := dao.GetUserByEmail(email)
	if err != nil {
		t.Fatalf("get user: %v", err)
	}
	return got
}

func seedProduct(t *testing.T, name string, price float64, stock int) uint {
	t.Helper()
	p := dao.Product{Name: name, Price: price, Stock: stock}
	if err := database.DB.Create(&p).Error; err != nil {
		t.Fatalf("seed product: %v", err)
	}
	return p.ID
}

func TestCartDAO_GetOrCreate_Add_Remove_Clear(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "cart@test.com")
	pid := seedProduct(t, "Zapa X", 99, 10)

	// get or create
	c1, err := dao.GetOrCreateCartByUserID(u.ID)
	if err != nil {
		t.Fatalf("GetOrCreateCartByUserID: %v", err) // <— faltaban args
	}
	if c1 == nil {
		t.Fatalf("carrito nil")
	}

	// add
	if err := dao.AddToCart(u.ID, pid, 2); err != nil {
		t.Fatalf("AddToCart err: %v", err) // <— faltaban args
	}

	// remove
	if err := dao.RemoveFromCart(u.ID, pid); err != nil {
		t.Fatalf("RemoveFromCart err: %v", err) // <— faltaban args
	}

	// clear (debería no fallar aunque esté vacío)
	if err := dao.ClearCart(u.ID); err != nil {
		t.Fatalf("ClearCart err: %v", err)
	}
}

func TestCartDAO_AddToCart_InvalidProduct_ReturnsError(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	u := seedUser(t, "bad@test.com")

	if err := dao.AddToCart(u.ID, 99999, 1); err == nil {
		t.Fatalf("esperaba error por product_id inválido")
	}
}

func TestCartDAO_AddTwice_IncrementsQuantity(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "inc@test.com")
	pid := seedProduct(t, "Zapa", 100, 50)

	if err := dao.AddToCart(u.ID, pid, 2); err != nil {
		t.Fatalf("AddToCart 1: %v", err)
	}
	if err := dao.AddToCart(u.ID, pid, 3); err != nil {
		t.Fatalf("AddToCart 2: %v", err)
	}

	cart, err := dao.GetOrCreateCartByUserID(u.ID)
	if err != nil {
		t.Fatalf("get cart: %v", err)
	}
	var item dao.CartItemEntity
	if err := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, pid).First(&item).Error; err != nil {
		t.Fatalf("query item: %v", err)
	}
	if item.Quantity != 5 {
		t.Fatalf("esperaba quantity=5, got %d", item.Quantity)
	}
}

func TestCartDAO_RemoveFromCart_MissingItem_ReturnsError(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "rm@test.com")
	pid := seedProduct(t, "Z", 10, 1)

	// no añadimos nada => eliminar debe devolver error
	if err := dao.RemoveFromCart(u.ID, pid); err == nil {
		t.Fatalf("esperaba error al eliminar item inexistente")
	}
}

func TestCartDAO_ClearCart_LeavesZeroItems(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "clear@test.com")
	pA := seedProduct(t, "A", 10, 10)
	pB := seedProduct(t, "B", 20, 10)

	if err := dao.AddToCart(u.ID, pA, 1); err != nil {
		t.Fatalf("add A: %v", err)
	}
	if err := dao.AddToCart(u.ID, pB, 2); err != nil {
		t.Fatalf("add B: %v", err)
	}
	if err := dao.ClearCart(u.ID); err != nil {
		t.Fatalf("clear: %v", err)
	}

	cart, err := dao.GetOrCreateCartByUserID(u.ID)
	if err != nil {
		t.Fatalf("get cart: %v", err)
	}
	var count int64
	if err := database.DB.Model(&dao.CartItemEntity{}).Where("cart_id = ?", cart.ID).Count(&count).Error; err != nil {
		t.Fatalf("count items: %v", err)
	}
	if count != 0 {
		t.Fatalf("esperaba 0 items en carrito, got %d", count)
	}
}
func TestCartDAO_GetOrCreate_CreatesNew(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	u := seedUser(t, "newcart@test.com")

	c, err := dao.GetOrCreateCartByUserID(u.ID)
	if err != nil {
		t.Fatalf("GetOrCreateCartByUserID: %v", err)
	}
	if c == nil || c.ID == 0 {
		t.Fatalf("esperaba carrito creado")
	}
}
