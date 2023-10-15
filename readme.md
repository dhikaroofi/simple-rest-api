# Gandiwa

This is a straightforward Golang application following the hexagonal architecture pattern with some custom modifications. Here's an overview of the project structure:
- bin -> Contains executable binary files.
- database/migrations -> Includes database migration scripts.
- pkg -> Holds custom code and initializes third-party libraries.
- logs -> Is used for logging files
- resources -> Stores all necessary resources for the application.
- internal ->
  - cmd -> Manages service aggregation, including the presentation layer and core use cases.
  - config -> Contains configuration settings.
  - common -> Houses utility functions.
  - presentation -> Handles communication with users, such as REST APIs.
  - usecase -> Contains the core application logic.

## Prerequisites

Before you get started, make sure you have the following prerequisites installed on your machine:

- Docker & Docker Compose: [Docker Installation Guide](https://docs.docker.com/get-docker/)
- Go (optional for building the application from source)

## Running the Application

To run the application using Docker, follow these steps:

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/your-golang-app.git
   cd your-golang-app


2. Set up the config
- rename '.env.example' to '.env' and make sure the values in it match your environment.
- replace './resources/config/config.yaml.example' with '/resources/config/config.yaml' and make sure the values are suitable for your environment.


3. build and run the apps with docker use this command below:

   ```bash
   make buildAndRunDockerApps
4. The app is ready to use. You can test it using Postman or any other HTTP client. To try the API, you can refer to the [API documentation](resources/docs/open-api.yaml) as a guideline."

 