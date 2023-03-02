Detail work you've completed in Sprint 2

List unit tests and Cypress test for frontend

List unit tests for backend

Add documentation for your backend API 

# Frontend




# Backend
## Documentation
Within the backend API, there are two functions that are used to "GET" or "POST" a user for the database.


**GetUser(w http.ResponseWriter, r *http.Request)** is used to get the user from the database, returning the username of the user


**PostUser(w http.ResponseWriter, r *http.Request)** is used to post the user to the database. It currently requires Username, Password, and Name to be provided so that the user will be posted.

These functions will be implemented to the itinerary struct in the future so the frontend can create or get a certain itinerary dependent on the itineraries a user owns.

## Unit Tests
We created unit tests for both PostUser and GetUser functions, however we find an issue with both functions resulting in a "FAIL" with the error message "invalid memory address or nil pointer dereference." This will be fixed in the future.
