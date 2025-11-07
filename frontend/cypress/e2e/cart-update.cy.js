// Sube la cantidad agregando el mismo producto de nuevo y verifica que el total aumenta
describe('Carrito - Actualizar cantidad (home -> /products -> carrito)', () => {
  const getCartTotal = () => {
    return cy.get('body').then(($b) => {
      const hasTestId = $b.find('[data-testid="cart-total"]').length > 0;
      const chain = hasTestId
        ? cy.get('[data-testid="cart-total"]', { timeout: 15000 })
        : cy.contains('p,div,span,strong,b,h2', /Total\s*:\s*\$/i, { timeout: 15000 });

      return chain
        .scrollIntoView({ block: 'center' })
        .invoke('text')
        .then((txt) => {
          const m = txt.match(/\$([\d.,]+)/);
          const num = m ? m[1] : '0';
          const normalized = num.replace(/\./g, '').replace(',', '.');
          return parseFloat(normalized);
        });
    });
  };

  beforeEach(() => {
    cy.viewport(1280, 900);
    cy.clearLocalStorage();
    cy.loginByApi();
  });

  it('agrega otra unidad del mismo producto y el total aumenta', () => {
    // 1) Home -> Ver Catálogo -> /products
    cy.visit('/');
    cy.contains('a,button', 'Ver Catálogo', { matchCase: false })
      .should('be.visible')
      .click();

    cy.url().should('include', '/products');

    // 2) Agregar un producto
    cy.contains('button', 'Agregar al carrito', { timeout: 15000 })
      .first()
      .scrollIntoView()
      .click({ force: true });

    // 3) Ir al carrito y leer total (ANTES)
    cy.get('[data-testid="navbar-cart"], a[href="/carrito"]', { timeout: 15000 })
      .first()
      .click();
    cy.url().should('include', '/carrito');

    getCartTotal().then((totalAntes) => {
      // 4) Volver a home y repetir a /products por el CTA (evitamos deep-link 404)
      cy.visit('/');
      cy.contains('a,button', 'Ver Catálogo', { matchCase: false })
        .should('be.visible')
        .click();
      cy.url().should('include', '/products');

      // 5) Agregar nuevamente el primer producto
      cy.contains('button', 'Agregar al carrito', { timeout: 15000 })
        .first()
        .scrollIntoView()
        .click({ force: true });

      // 6) Volver al carrito y comparar totales
      cy.get('[data-testid="navbar-cart"], a[href="/carrito"]', { timeout: 15000 })
        .first()
        .click();
      cy.url().should('include', '/carrito');

      getCartTotal().then((totalDespues) => {
        expect(totalDespues).to.be.greaterThan(totalAntes);
      });
    });
  });
});
