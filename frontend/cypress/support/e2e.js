// ***********************************************************
// This example support/e2e.js is processed and
// loaded automatically before your test files.
//
// This is a great place to put global configuration and
// behavior that modifies Cypress.
//
// You can change the location of this file or turn off
// automatically serving support files with the
// 'supportFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/configuration
// ***********************************************************

// Import commands.js using ES2015 syntax:
import './commands'
// Comando para login por API y dejar token en localStorage

// frontend/cypress/support/e2e.js

Cypress.on('uncaught:exception', (err, runnable) => {
  // Si querés ser específico, sólo ignoramos el famoso "Unexpected token '<'"
  if (err.message && err.message.includes("Unexpected token '<'")) {
    // returning false prevents Cypress from failing the test
    return false;
  }

  // para otros errores, seguí fallando como siempre
  return true;
});


Cypress.Commands.add('loginByApi', () => {
  const api = 'https://backendqa-production.up.railway.app';

  const email = Cypress.env('email') || 'simon@example.com';
  const password = Cypress.env('password') || 'simon123';

  return cy.request({
    method: 'POST',
    url: `${api}/api/login`,     // 👈 ajustá si tu endpoint es otro
    body: { email, password },        // 👈 keys según tu backend
    failOnStatusCode: false,          // para ver la respuesta aunque sea 4xx/5xx
  }).then((res) => {
    // Debug útil si algo falla:
    cy.log(`Login status: ${res.status}`);
    // Esperamos 200/201 del login:
    expect(res.status).to.be.oneOf([200, 201]);

    const token = res.body?.token || res.body?.data?.token;
    const user = res.body?.user || res.body?.data?.user || {};

    // Guarda como lo hace tu frontend:
    window.localStorage.setItem('token', token);
    window.localStorage.setItem('user', JSON.stringify(user));
  });
});
