// cypress/e2e/cart-error.cy.js
// Valida integración FE–BE sin token: navegar vía SPA y comprobar UI resultante
describe('Carrito - Manejo de errores sin token', () => {
  beforeEach(() => {
    cy.viewport(1280, 900);
    cy.clearLocalStorage(); // sin token
  });

  it('desde Home abre Carrito por SPA y muestra mensaje o carrito vacío', () => {
    // 1) Home
    cy.visit('/');

    // 2) Abrir carrito desde el navbar (SPA: evita 404 del deep-link)
    cy.get('[data-testid="navbar-cart"], a[href="/carrito"]', { timeout: 15000 })
      .first()
      .click();

    cy.url().should('include', '/carrito');

    // 3) Dos posibles UIs según tu componente:
    //    a) Mensaje de auth requerida (si no hace early return)
    //    b) "Tu carrito está vacío" (early return cuando no hay items)
    cy.get('body').then(($b) => {
      const hasMsg = /debes iniciar sesión para ver el carrito/i.test($b.text());
      const hasEmpty = /tu carrito está vacío/i.test($b.text());

      expect(hasMsg || hasEmpty, 'mensaje o estado vacío').to.be.true;

      if (hasMsg) {
        cy.contains(/debes iniciar sesión para ver el carrito/i).should('be.visible');
      } else {
        cy.contains(/tu carrito está vacío/i).should('be.visible');
        // y el CTA a productos
        cy.contains('a', /ver productos/i).should('be.visible');
      }
    });
  });
});
