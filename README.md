# User web Service

This web service has basic endpoints for working with user data, such as creating, deleting and getting a user, and also implements a system of subscriptions and unsubscribes. The project is an example of the use of GraphQL in Go.

## Technologies Used

- **Language**: `Go`
- **API's**: `REST`, `gRPC`, `GraphQL`
- **Technologies**: `GraphQL`, `Docker`, `Docker Compose`, `mockery`, `testify`, `protobuff`, `Swagger`
- **Storage**: `MongoDB`

## Additional Notes

- In the project directory there is a client directory, that allows you to run gRPC & GraphQL calls to the server API.

## Endpoints
You can view the documentation by starting the server and accessing endpoint `/docs`

## Setup Instructions
1. To run the application and set up mongoDB:

   ```bash
   make docker-compose-up
   ```

2. To run tests with coverage and generate page with coverage
   ```bash
   make test
   ```