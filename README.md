# go-recipe-api

The go-recipe-api is the final form of my Online Cookbook Backend, built to be hosted on local hardware utilizing Docker. It is built using Go, Gin and Gorm.

## To-Do

- [ ] Work out image handling function
- [ ] Work out permanent email handling solution
- [ ] Limit registrations to local (new users must verify before being able to login, which can only be done locally for now)
- [ ] TBD

## Getting Started

Setup a PostgreSQL database and add the appropriate environment variables to the `app.env` file. Follow the `app.env.example` file for guidance.

Next, run the following commands to start the server:

```bash
From the root directory, run the following commands:

```bash
go mod init github.com/W5DEV/go-recipe-api

go run migrate/migrate.go

air
```
