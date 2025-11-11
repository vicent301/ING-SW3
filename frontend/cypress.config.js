const { defineConfig } = require('cypress');

module.exports = defineConfig({
  e2e: {
    // Usa directamente la URL QA si no se define CYPRESS_BASE_URL
    baseUrl: process.env.CYPRESS_BASE_URL || 'https://cont-front-qa-dvctbgbdfbbrfpbe.chilecentral-01.azurewebsites.net',
    
    specPattern: 'cypress/e2e/**/*.cy.{js,jsx,ts,tsx}',
    supportFile: 'cypress/support/e2e.js',
    video: true,
    screenshotsFolder: 'cypress/screenshots',
    videosFolder: 'cypress/videos',

    setupNodeEvents(on, config) {
      // Podés usar hooks si más adelante integrás reporters como Allure
      return config;
    },
  },

  // Reporte JUnit para Azure DevOps
  reporter: 'junit',
  reporterOptions: {
    mochaFile: 'cypress/results/junit-[hash].xml',
    toConsole: false,
  },
});
