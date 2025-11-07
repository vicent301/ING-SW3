describe('Carrito - Actualizar cantidad', () => {
  beforeEach(() => {
    cy.viewport(1280, 900);
    cy.clearLocalStorage();
    cy.loginByApi();
    cy.visit('/');

    // Si el carrito puede estar vacío, mete 1 producto rápido:
    cy.contains('a,button', 'Ver Catálogo', { matchCase: false }).click();
    cy.get('[data-testid="product-card"], .product-card').first().within(() => {
      cy.contains('button', 'Agregar al carrito', { matchCase: false }).click();
    });
    cy.get('[data-testid="navbar-cart"], a[href="/carrito"]').first().click();
    cy.url().should('include', '/carrito');
  });

  it('incrementa cantidad y actualiza el total', () => {
    cy.get('[data-testid="cart-total"], .cart-total')
      .invoke('text')
      .then((totalAntes) => {
        cy.get('[data-testid="cart-qty-inc"], .qty-inc, button:contains("+")').first().click();
        cy.wait(800);
        cy.get('[data-testid="cart-total"], .cart-total').should(($t) => {
          expect($t.text().trim()).not.to.eq(totalAntes.trim());
        });
      });
  });
});
