## Work Completed

**Josh Lamb** - Finished implementing status function, deleted a generic "get" HTTP route for a user and replaced it with a route to register a user, replaced all route query parameters with route parameters. Implemented bcrypt to store password hashes on the database instead of the passwords themselves, fixed Go unit testing functions for the HTTP routes, started looking into JWT tokens with Gabe.

**Gabriel Cortez** - Mostly finished the itinerary class to store the information needed for a created itinerary, with a map that will hold any locations that have been voted on. Created some basic tests for "PUT" and "GET" for itinerary, but have not been able to test yet. Created a simple JWT token template to be used in the future but may need to change some parts of the code in other files in the future.

**Tony Gupta** - Reworked all of the logic behind login and signup (component .ts files, auth-service, auth-guard) to actually connect to backend on localhost9000 and send POST requests. Sends a JSON containing user data to either add to database or check against database and obtain a JWT for login. Some issue with CORS (requires more research) that denied backend access from the frontend's localhost4200 location. Will also need to store JWT in a cookie instead of localstorage and implement an HTTPinterceptor.

**Lucas Mueller** - Wrote more cypress tests to imporove the specificity of our testing and overall quality of our app. The cypress tests included verifying the page loads correctly by checking the HTTP status code, checking to make sure all the elements are visible as expected, and chcking how the app will display on different screen sizes. Also, researched integration methods for backend and frontend communication. JWT tokens is the authentication method we aim to use for login so I also looked into that. the goal for next sprint is to have the backend and frontend communicate using HTTP requests.


## Frontend Unit Tests and Cypress Tests

we added three basic tests:

test 1: verifies that the page loads successfully and has the expected title and status code.
test 2: verifies that the page contains the expected content and that a specific element is visible and has the expected text.
test 3: verifies that the page is responsive on different screen sizes.

## Backend Unit Tests
We fixed the unit test for the RegisterUser function from the "/user/register" route. We removed the unit test for the GetUser function from the "/user/get/{username}" route which we deleted. Finally, we added a new unit test for our new Login function from the new "/user/login" route. 

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

There are currently 2 additional routes that are mostly written, but not fully implemented or tested.

**4.** /itinerary/get/{id}
description: attempts to find a specific itinerary from the database
- requires: `[id: <int>]`
- returns: the itinerary's full database entry, including the ID, name, address, radius, and saved locations

**5.** /itinerary/post/{name}/{address}/{radius}
- description: creates an itinerary in the database
- requires: `[name: <string>, address: <string>, radius: <string>]`
- returns: none
