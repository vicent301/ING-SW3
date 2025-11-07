describe('Carrito - Crear ítem (via Ver Catálogo)', () => {
  beforeEach(() => {
    cy.viewport(1280, 900);
    cy.clearLocalStorage();
    cy.loginByApi();
    cy.visit('/');
  });

  it('abre catálogo y agrega un producto al carrito', () => {
    // 1) Ir al catálogo
    cy.contains('a,button', 'Ver Catálogo', { matchCase: false })
      .should('be.visible')
      .click();

    // 2) Agregar un producto
    cy.contains('button', 'Agregar al carrito', { timeout: 10000 })
      .first()
      .scrollIntoView()
      .click({ force: true });

    // 3) Ir al carrito
    cy.get('[data-testid="navbar-cart"], a[href="/carrito"]', { timeout: 10000 })
      .first()
      .click();

    // 4) Confirmar URL
    cy.url().should('include', '/carrito');

    // 5) Verificar TOTAL (con o sin data-testid)
    cy.get('body').then(($b) => {
      if ($b.find('[data-testid="cart-total"]').length) {
        // Caso con testid (build nuevo)
        cy.get('[data-testid="cart-total"]', { timeout: 15000 })
          .scrollIntoView({ block: 'center' })
          .should('be.visible')
          .invoke('text')
          .should('match', /Total\s*:\s*\$\d/);
      } else {
        // Caso sin testid (QA viejo o caché): buscar el texto "Total: $"
        cy.contains('p,div,span,strong,b,h2', /Total\s*:\s*\$/i, { timeout: 15000 })
          .scrollIntoView({ block: 'center' })
          .should('be.visible')
          .invoke('text')
          .should('match', /Total\s*:\s*\$\d/);
      }
    });
  });
});
