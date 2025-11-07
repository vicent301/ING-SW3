describe('Errores de integración', () => {
  it('muestra mensaje amigable si falla /api/cart', () => {
    // Fuerza error del backend para este test
    cy.intercept('GET', '/api/cart*', { statusCode: 500, body: { message: 'Server error' } }).as('cartFail')

    cy.visit('/carrito')
    cy.wait('@cartFail')
    cy.get('[data-testid="alert-error"]').should('contain.text', 'No se pudo cargar el carrito')
  })
})
