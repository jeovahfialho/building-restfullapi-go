
# Book Management API

This project is a simple RESTful API for book management. It allows users to create, read, update, and delete books from a database. The API is built with Go and uses MySQL for data persistence.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Docker
- Docker Compose

### Installing

1. Clone the repository to your local machine:

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

This command will build the Go application and MySQL database containers. The database will be automatically set up with the necessary tables.

## API Endpoints

Below are the available API endpoints with examples:

### Create a New Book

- **POST** `/books`

```json
// Request Body
{
  "title": "The Name of the Wind",
  "author": "Patrick Rothfuss",
  "publishedYear": 2007,
  "genre": "Fantasy",
  "summary": "The tale of the magically gifted young man who grows to be the most notorious wizard his world has ever seen."
}

// Response
Status: 201 Created
```

### Get All Books

- **GET** `/books`

```json
// Response
Status: 200 OK
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

```json
// Response
Status: 200 OK
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

```json
// Request Body
{
  "title": "The Wise Man's Fear",
  "author": "Patrick Rothfuss",
  "publishedYear": 2011,
  "genre": "Fantasy",
  "summary": "Sequel to The Name of the Wind, continuing the story of Kvothe..."
}

// Response
Status: 200 OK
```

### Delete a Book

- **DELETE** `/books/{id}`

```json
// Response
Status: 200 OK
{
  "message": "Book deleted successfully"
}
```

## Running the Tests

Explain how to run the automated tests for this system.

(Add instructions based on your project's test setup)

## Built With

- [Go](https://golang.org/) - The Go Programming Language
- [MySQL](https://www.mysql.com/) - The MySQL Database
- [Docker](https://www.docker.com/) - Containerization Platform

## Authors

- **Your Name** - *Initial work* - [YourGitHub](https://github.com/YourGitHub)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
