REST API Todo app

For this application you need download some dependencies
go get github.com/jmoiron/sqlx - working with PostgreSQl
go get github.com/lib/pq - driver for PostgreSQL
go get github.com/spf13/viper - configuration solution
go get github.com/joho/godotenv - configuration solution
go get -u github.com/golang-jwt/jwt/v5 - library for creating and parsing JWT 

Also, you need to run Docker container with PostgreSQL
docker pull postgres
sudo docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres

Launch migration in project directory
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

In this app, you can create users, create list and items. There also JWT autentication for signed in users.
