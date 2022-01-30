# Final Project Sanbercode Go Batch 31

## Tech Stack
- Go 1.17
- Gin 1.7.7
- Postgresql 9.6.x

## Database Schema
- [Database Schema](/erd.png)

## Endpoint Documentation
- https://blog-api-fp.herokuapp.com/api/v1/swagger/index.html

### Installation (Go)
- Copy .env.example and rename to .env
- Set your own configuration in .env
- Run `go get`
- Run `go run main.go`

### Installation (Docker)
- Copy .env.example and rename to .env
- Set in .env
  ```
  DB_HOST=blog-api-db
  DB_PORT=5432
  DB_USER=blog
  DB_PASS=secret
  DB_NAME=blog_api
  ```
- Run `docker compose build`
- Run `docker compose up -d`
- Ready to rock and roll