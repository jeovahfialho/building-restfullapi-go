
# Book Management API

This project is a simple RESTful API for managing books. It supports operations to create, read, update, and delete (CRUD) books from a MySQL database. The API is developed in Go and designed to be containerized with Docker for easy development, testing, and deployment.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Postman](https://www.postman.com/) (for API testing)

### Installing

1. Clone the repository:
   ```bash
   git clone https://github.com/jeovahfialho/building-restfullapi-go.git
   ```

2. Navigate to the project directory:
   ```bash
   cd building-restfullapi-go
   ```

3. Use Docker Compose to build and start the services:
   ```bash
   docker-compose up --build
   ```

## API Endpoints

The following API endpoints are available:

### Create a New Book

- **POST** `/books`
  - **Request Body Example**:
    ```json
    {
      "title": "The Name of the Wind",
      "author": "Patrick Rothfuss",
      "publishedYear": 2007,
      "genre": "Fantasy",
      "summary": "The tale of the magically gifted young man..."
    }
    ```
  - **Response**: `201 Created`

### Get All Books

- **GET** `/books`
  - **Response**: `200 OK`
    ```json
    [
      {
        "id": 1,
        "title": "The Name of the Wind",
        "author": "Patrick Rothfuss",
        "publishedYear": 2007,
        "genre": "Fantasy",
        "summary": "The tale of the magically gifted young man..."
      }
    ]
    ```

### Get a Specific Book

- **GET** `/books/{id}`
  - **Response**: `200 OK`
    ```json
    {
      "id": 1,
      "title": "The Name of the Wind",
      "author": "Patrick Rothfuss",
      "publishedYear": 2007,
      "genre": "Fantasy",
      "summary": "The tale of the magically gifted young man..."
    }
    ```

### Update a Book

- **PUT** `/books/{id}`
  - **Request Body Example**:
    ```json
    {
      "title": "The Wise Man's Fear",
      "author": "Patrick Rothfuss",
      "publishedYear": 2011,
      "genre": "Fantasy",
      "summary": "Sequel to The Name of the Wind..."
    }
    ```
  - **Response**: `200 OK`

### Delete a Book

- **DELETE** `/books/{id}`
  - **Response**: `200 OK`
    ```json
    {
      "message": "Book deleted successfully"
    }
    ```

## Running the Tests

To run the automated tests for this system, follow these steps:

1. Ensure the Docker containers for the application and database are running.
2. Execute the test command:
   ```bash
   docker-compose exec app go test -v ./...
   ```

This will run all the automated tests defined in the project, outputting the results to your terminal.

## Postman Collection

To facilitate API testing, a Postman collection is included with predefined requests for all available endpoints. Import the following collection into your Postman application to get started:

[Book Management Postman Collection](https://martian-shuttle-401403.postman.co/workspace/Tests---Jeovah~c7ebacd1-af72-4324-b67f-86b96865fc22/collection/1780378-517b950f-d054-4a89-a1cf-8b03e5f494b9?action=share&source=collection_link&creator=1780378)

## Built With

- [Go](https://golang.org/) - The Go Programming Language
- [MySQL](https://www.mysql.com/) - The MySQL Database
- [Docker](https://www.docker.com/) - Containerization Platform