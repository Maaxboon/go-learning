# Agenda

- Request/Response Cycle
- What is an API
- How to use the fetch method

# Build a Single Page Application (SPA) random fact generator that:

- Makes a GET request to an API
- Manipulates the DOM with the response from the request

# API

- API stands for Application programming interface.
- It is a software with protocols and procedures to send data back and forth between to separate computers i.e a web page or app and a user.

# Request/Response Cycle

- Client makes a request to look at a friends photo.
- The type of request is defined by the HTTP verb included with the request.
- Most common HTTP verbs: GET, POST, PUT/PATCH, DELETE
  - GET: I just want information
  - POST: Send some type of data. I want to add my own picture, liking,, commenting.
  - PUT/PATCH: Want to edit. Uploaded wrong picture and want to update it. Twitter edits, WhatsApp photo edits.
  - DELETE: Remove image, comment, like, text
- The server receives and processes the request.
- Communication between server and database.
- Then sends response back to the client.
