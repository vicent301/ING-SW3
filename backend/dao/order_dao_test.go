package dao_test

import (
	"backend/dao"
	"backend/database"
	"backend/testutil"
	"testing"
)

func TestOrderDAO_CreateAndList(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "order@test.com")
	pid := seedProduct(t, "Zapa Y", 120, 5)

	// agrego al carrito
	if err := dao.AddToCart(u.ID, pid, 3); err != nil {
		t.Fatalf("AddToCart: %v", err)
	}

	// creo la orden
	o, err := dao.CreateOrderFromCart(u.ID)
	if err != nil {
		t.Fatalf("CreateOrderFromCart: %v", err)
	}
	if o == nil {
		t.Fatalf("orden nil")
	}

	// listado
	list, err := dao.GetOrdersByUser(u.ID)
	if err != nil {
		t.Fatalf("GetOrdersByUser: %v", err)
	}
	if len(list) == 0 {
		t.Fatalf("esperaba al menos 1 orden")
	}
}

func TestOrderDAO_CreateOrder_EmptyCart_ReturnsError(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	u := seedUser(t, "nocart@test.com")

	if _, err := dao.CreateOrderFromCart(u.ID); err == nil {
		t.Fatalf("esperaba error por carrito vacío")
	}
}

func TestOrderDAO_CreateOrder_ClearsCart(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "ordclear@test.com")
	p := seedProduct(t, "Zapa", 50, 10)

	if err := dao.AddToCart(u.ID, p, 2); err != nil {
		t.Fatalf("add: %v", err)
	}
	if _, err := dao.CreateOrderFromCart(u.ID); err != nil {
		t.Fatalf("create order: %v", err)
	}

	cart, err := dao.GetOrCreateCartByUserID(u.ID)
	if err != nil {
		t.Fatalf("get cart: %v", err)
	}
	var count int64
	if err := database.DB.Model(&dao.CartItemEntity{}).Where("cart_id = ?", cart.ID).Count(&count).Error; err != nil {
		t.Fatalf("count: %v", err)
	}
	if count != 0 {
		t.Fatalf("esperaba carrito vacío post-orden, got %d items", count)
	}
}

func TestOrderDAO_ListAfterCreate_ReturnsOrder(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	u := seedUser(t, "ordlist@test.com")
	p := seedProduct(t, "Prod", 30, 9)

	if err := dao.AddToCart(u.ID, p, 1); err != nil {
		t.Fatalf("add: %v", err)
	}
	if _, err := dao.CreateOrderFromCart(u.ID); err != nil {
		t.Fatalf("create: %v", err)
	}
	orders, err := dao.GetOrdersByUser(u.ID)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(orders) == 0 {
		t.Fatalf("esperaba al menos 1 orden")
	}
}
