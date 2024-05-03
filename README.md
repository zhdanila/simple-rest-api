# REST API Todo App

This is a Todo application that provides RESTful APIs for managing users, lists, and items. It also includes JWT authentication for signed-in users.

## Dependencies

Before running the application, make sure to download the following dependencies:

- [github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx): Working with PostgreSQL
  ```
  go get github.com/jmoiron/sqlx
  ```

- [github.com/lib/pq](https://github.com/lib/pq): PostgreSQL driver
  ```
  go get github.com/lib/pq
  ```

- [github.com/spf13/viper](https://github.com/spf13/viper): Configuration solution
  ```
  go get github.com/spf13/viper
  ```

- [github.com/joho/godotenv](https://github.com/joho/godotenv): Configuration solution
  ```
  go get github.com/joho/godotenv
  ```

- [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt): Library for creating and parsing JWT tokens
  ```
  go get -u github.com/golang-jwt/jwt/v5
  ```

## Setting up PostgreSQL with Docker

To run the application, you need to set up a PostgreSQL database using Docker:

1. Pull the PostgreSQL Docker image:
   ```
   docker pull postgres
   ```

2. Run the Docker container with PostgreSQL:
   ```
   sudo docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
   ```

## Running Migrations

After setting up PostgreSQL, you need to run migrations in the project directory to create the necessary database schema:

```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

## Usage

Once dependencies are installed and migrations are applied, you can start using the Todo app to create users, lists, and items. JWT authentication is provided for signed-in users.
