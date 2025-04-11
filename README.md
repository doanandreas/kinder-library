# 📚 Kinder Library API

A simple RESTful API to manage books in a library system using Go.

---

## 📋 Prerequisites

Ensure you have the following installed:

- **Docker & Docker Desktop**: [https://www.docker.com/get-started/](https://www.docker.com/get-started/)

---

## 🚀 How to run the API

1. Create a `.env` file using the contents of `.env.example`.

2. Run using Docker Compose:

```bash
docker compose up --build
```

3. Once running, access the API at:

```
http://localhost:8080/v1/books
```

---

## ✅ How to run tests

Run all tests inside a Docker container:

```bash
docker build -f Test.Dockerfile -t kinder-library-test .
docker run --rm kinder-library-test
```

---

## 📖 How to Check API Docs

The Swagger UI is embedded. View the API documentation by:

1. Run the application with `docker compose up --build`, and
2. Open API docs at `http://localhost:8080/v1/swagger`.

---

## 🔐 Authentication

All endpoints (except `/v1/healthcheck`) require **Basic Authentication**.

### Hardcoded Credentials:

- **Username:** `doan`
- **Password:** `didinding`

### Example Header:

```
Authorization: Basic ZG9hbjpkaWRpbmRpbmc=
```

(This is Base64-encoded `doan:didinding`)

---

## 🌐 API Endpoints Overview

| Method | Endpoint            | Description              | Auth Required |
|--------|---------------------|--------------------------|----------------|
| GET    | `/v1/healthcheck`   | Check API availability   | ❌             |
| GET    | `/v1/books`         | List all books           | ✅             |
| POST   | `/v1/books`         | Insert a new book        | ✅             |
| PUT    | `/v1/books/{id}`    | Update an existing book  | ✅             |
| DELETE | `/v1/books/{id}`    | Delete a book            | ✅             |

---

## 📂 Project Structure

```
.
├── cmd/api              # Entry point and HTTP handlers
├── internal/data        # Data models and validation logic
├── internal/mocks       # Mock files for testing
├── internal/repository  # DB interaction logic
├── internal/validator   # Input validation helpers
├── go.mod / go.sum      # Go modules
├── .env.example         # Example .env file. Copy and rename it to .env before running.
├── docker-compose.yaml  # Docker Compose configuration files
├── Dockerfile           # Dockerfile for building API binaries
├── Test.Dockerfile      # Dockerfile for running tests
├── wait-for-it.sh       # Script to wait for Postgres container
```

---

## 👤 Author

Made with 💛 by Doan Andreas Nathanael for KinderCastle take-home interview.
