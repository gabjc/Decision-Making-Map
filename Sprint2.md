## Work Completed

**Josh Lamb** - combined README.md and README.txt, replaced pure SQL with GORM, added how to auto-install Go packages to README, created Go unit test for "user/get" route

**Gabriel Cortez** - tried to figure out how to implement a way to store the multiple variables that a user owns within a user's struct, developed the itinerary struct. Made the unit test for PostUser.

**Tony Gupta** - worked on html for pages (used material to create forms for login and register), created public and protected angular modules to implement later functionality preventing users who aren't logged in to access the main app content, buggy ts button logic to send users to dashboard.

**Lucas Mueller** - created the system for routing our website between multiple pages, wrote cypress tests to make sure the features we have implemented so far are behaving as expected.


## Frontend Unit Tests and Cypress Tests

To test the fron end, we wrote a few tests that check to make sure the angular components we are using exist and have all the right values. In addition, we wrote tests that check that the routing between the different pages is functioning properly.

## Backend Unit Tests

We created unit tests for both PostUser and GetUser functions, which implement the "user/post" and "user/get" routes respectively, however a current issue is causing both functions to result in a "FAIL" with the error message "invalid memory address or nil pointer dereference." This will be fixed in the future.

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
- returns: status code 200

**4.** /itinerary/get/{id}
description: attempts to find a specific itinerary from the database
- requires: `[id: <int>]`
- returns: the itinerary's full database entry, including the ID, name, address, radius, and saved locations

**5.** /itinerary/post/{name}/{address}/{radius}
- description: creates an itinerary in the database
- requires: `[name: <string>, address: <string>, radius: <string>]`
- returns: none
