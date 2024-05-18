# go-recipe-api

The go-recipe-api is the final form of my Online Cookbook Backend, built to be hosted on local hardware utilizing Docker. It is built using Go, Gin and Gorm.

## To-Do

- [ ] Work out image handling functionality - Dockerized image storage??
- [ ] Work out permanent solution for handling emails
- [ ] Create superuser ability to invite new users (send email with validation link to register, set password, name, etc)
- [ ] Combine PostgreSQL configuration in docker to link setup (create if not exists)

## Getting Started

Setup a PostgreSQL database and add the appropriate environment variables to the `app.env` file. Follow the `app.env.example` file for guidance.

Next, run the following commands from the root directory to start the server:

```bash
go mod init github.com/W5DEV/go-recipe-api

go run migrate/migrate.go

air
```
