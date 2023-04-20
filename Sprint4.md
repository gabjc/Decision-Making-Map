## Work Completed

**Josh Lamb** - Fixed CORS issues, implemented bcrypt for hashing passwords to store on database, implmeneted JWT token to send to the frontend, assisted Gabe with troubleshooting unit tests, assisted Tony with troubleshooting navigation between webpages on the frontend.

**Gabriel Cortez** - Finished the itinerary struct and the functions with it, including a create and get function for the itineraries. Created some tests and got help from Josh to fix the tests. Set up the functions for the JWT token.

**Tony Gupta** - 

**Lucas Mueller** - 


## Frontend Unit Tests and Cypress Tests

we added three basic tests:



## Backend Unit Tests
We have a unit test ofr RegisterUser that passes, and a test for Login that also passes. As for the itinerary tests, there is the test to CreateItinerary, which passes. However, for GetAllItineraries it says fail, but this is due to a problem with the formatting from switching a string to json and vice versa.

## Backend API Documentation
This API assumes that all responses are made with `Content-Type: application/json`
There are currently 3 functioning routes to interact with the backend, specifically for registering a new user, logging in, and getting a generic status OK to confirm front-end and back-end communication.

**1.** /user/register
- description: creates a user in the database
- requires: JSON taken by register route -> `{"name": "namenamename", "email": "email@email.com", "hash": "myPassword"}`
- returns: returns 201 status Created to indicate user created

**2.** /user/login
- description: checkes to see if the passed in email and password match to a pair in the database
- requires: JSON taken by login route -> `{"email": "email@email.com", "hash": "myPassword"}` 
- returns: returns 404 status Not Found if user doesnt exist, 403 status Forbidden if password is wrong
NOTE: we will return 201 status Created when we fully implement the JWT token with the login

**3.** /status
- description: checks the status of the backend
- requires: `none`
- returns: status code 200 OK to indicate connection between frontend and backend

**4.** /itinerary/get/{id}
description: attempts to find all itineraries in the database that has the creatorID, which is the id of the user that created the itinerary.
- requires: `[creator_id: <int>]`
- returns: all the itineraries with the passed in creator_id, and thier full database entry, including the ID, name, plans, user, and the user's id

**5.** /itinerary/home
- description: creates an itinerary in the database
- requires: `N/A`
- returns: none
