describe('Basic Cypress Tests', () => {
    beforeEach(() => {
      // Load the page before each test
      cy.visit('http://localhost:4200/')
    })
  
    it('should load the page successfully', () => {
      // Verify that the page title is correct
      cy.title().should('eq', 'App')
  
      // Verify that the page has a 200 status code
      cy.request('http://localhost:4200/').its('status').should('eq', 200)
    })
  
    it('should display the expected content', () => {
      // Verify that the page contains the expected text
      cy.contains('Name Forthcoming')
  
      // Verify that a specific element is visible and contains the expected text
      //cy.get('button').should('be.visible').and('have.text', 'Login')
    })
  
    it('should be responsive on different screen sizes', () => {
      // Set the viewport to a mobile screen size
      cy.viewport('iphone-6')
  
      // Verify that the page content is still visible
      cy.contains('Name Forthcoming')
  
      // Set the viewport to a tablet screen size
      cy.viewport('ipad-2')
  
      // Verify that the page content is still visible
      cy.contains('Name Forthcoming')
  
      // Set the viewport to a desktop screen size
      cy.viewport('macbook-15')
  
      // Verify that the page content is still visible
      cy.contains('Name Forthcoming')
    })
  })
  