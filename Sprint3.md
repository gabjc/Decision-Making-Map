## Work Completed

**Josh Lamb** - Finished implementing status function, deleted a generic "get" HTTP route for a user and replaced it with a route to register a user, replaced all route query parameters with route parameters. Implemented bcrypt to store password hashes on the database instead of the passwords themselves, fixed Go unit testing functions for the HTTP routes, started looking into JWT tokens with Gabe.

**Gabriel Cortez** - Mostly finished the itinerary class to store the information needed for a created itinerary, with a map that will hold any locations that have been voted on. Created some basic tests for "PUT" and "GET" for itinerary, but have not been able to test yet. Created a simple JWT token template to be used in the future but may need to change some parts of the code in other files in the future.

**Tony Gupta** - 

**Lucas Mueller** - 


## Frontend Unit Tests and Cypress Tests



## Backend Unit Tests
We fixed the unit test for the RegisterUser function from the "/user/register" route. We removed the unit test for the GetUser function from the "/user/get/{username}" route which we deleted. Finally, we added a new unit test for our new Login function from the new "/user/login" route. 

## Backend API Documentation
This API assumes that all responses are made with `Content-Type: application/json`
There are currently 2 functioning routes to interact with the backend, specifically for creating a new user and finding a specific user.

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

There are currently 2 additional routes that are mostly written, but not fully implemented or tested.

**4.** /itinerary/get/{id}
description: attempts to find a specific itinerary from the database
- requires: `[id: <int>]`
- returns: the itinerary's full database entry, including the ID, name, address, radius, and saved locations

**5.** /itinerary/post/{name}/{address}/{radius}
- description: creates an itinerary in the database
- requires: `[name: <string>, address: <string>, radius: <string>]`
- returns: none
