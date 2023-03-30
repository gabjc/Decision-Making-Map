## Backend API Documentation
This API assumes that all responses are made with `Content-Type: application/json`
There are currently 2 functioning routes to interact with the backend, specifically for creating a new user and finding a specific user.

**1.** /user/get/{username}
- description: attempts to find a specific user from the database
- requires: `[username: <string>]`
- returns: the user's full database entry, including their username, password, and name on success, else a blank entry with the previous variables as empty strings on failure

**2.** /user/post/{username}/{password}/{name}
- description: creates a user in the database
- requires: `[username: <string>, password: <string>, name: <string>]`
- returns: none

There are currently 3 additional routes that are mostly written, but not fully implemented or tested.

**3.** /status
- description: checks the status of the backend
- requires: `none`
- returns: status code 200 OK to indicate connection between frontend and backend

**4.** /itinerary/get/{id}
description: attempts to find a specific itinerary from the database
- requires: `[id: <int>]`
- returns: the itinerary's full database entry, including the ID, name, address, radius, and saved locations

**5.** /itinerary/post/{name}/{address}/{radius}
- description: creates an itinerary in the database
- requires: `[name: <string>, address: <string>, radius: <string>]`
- returns: none
