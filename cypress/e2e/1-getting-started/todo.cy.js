/// <reference types="cypress" />

describe('HomePage Test', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/')
  })

  it('There should be 2 buttons', () => {
    
    cy.get('.card').should('have.length', 2)
    
    cy.get('.card').first().should('have.text', 'Login')
    cy.get('.card').last().should('have.text', 'Dont Have An Account?')
  })

  it('Route Testing Login Button', () => {
    cy.get('.card').first().click()
    cy.location("pathname").should("eq", "/login")
  })

  it('Route Testing Sign-In Button', () => {
    cy.get('.card').last().click()
    cy.location("pathname").should("eq", "/signin")
  })

})
