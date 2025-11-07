describe('Carrito - Crear ítem (via Ver Catálogo)', () => {
  beforeEach(() => {
    cy.viewport(1280, 900);
    cy.clearLocalStorage();
    cy.loginByApi();
    cy.visit('/');
  });

  it('abre catálogo y agrega un producto al carrito', () => {
    cy.contains('a,button', 'Ver Catálogo', { matchCase: false })
      .should('be.visible')
      .click();

    cy.contains('button', 'Agregar al carrito', { timeout: 10000 })
      .first()
      .scrollIntoView()
      .click({ force: true });

    cy.get('[data-testid="navbar-cart"], a[href="/carrito"]', { timeout: 10000 })
      .first()
      .click();

    cy.url().should('include', '/carrito');

    cy.get('[data-testid="cart-total"]', { timeout: 10000 })
      .scrollIntoView()
      .should('be.visible')
      .invoke('text')
      .should('match', /Total\s*:\s*\$\d/);
  });
});
