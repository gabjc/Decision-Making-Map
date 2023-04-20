## Work Completed

**Josh Lamb** - 

**Gabriel Cortez** - 

**Tony Gupta** - 

**Lucas Mueller** - 


## Frontend Unit Tests and Cypress Tests

we added three basic tests:



## Backend Unit Tests

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
description: attempts to find a specific itinerary from the database
- requires: `[id: <int>]`
- returns: the itinerary's full database entry, including the ID, name, address, radius, and saved locations

**5.** /itinerary/home
- description: creates an itinerary in the database
- requires: `[name: <string>, address: <string>, radius: <string>]`
- returns: none
