# Couchless Backend

Couchless is an open source application and backend server to self host your fitness and health data. Stop giving away control over your data.

The frontend app is also open source on Github: https://github.com/fusion44/couchless-frontend

This project is in its very early stages of development and may contain bugs, use with caution.

## Basic Setup

### Postgres

A working PostgresSQL instance is required. Update the .env and the config.toml files with the Postgres credentials.

### Fit2JSON

Fit2json is a tool to convert .FIT files to json files. The binary is called by the server to extract necessary information from the .FIT files.

Linux binaries can be found on Github: https://github.com/fusion44/fit2json/releases

Download the binary and edit the fit2json-path variable in config.toml with the path.

### Execute migrations

Execute these commands from within the root project folder or change the path's accordingly. Only migrate version v4.11.0 was tested thus far, but a newer version should work as well. At some point in the future this step will be redundant.

#### Get migrate executable

In project root:\
`curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz`

#### Migrate up

Make sure you have edited the `.env` file with the postgres instance credential before running this.\
`source .env && ./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL up`

### Run the server

In project root:\ `go run server.go`

Open GraphQL Playground: [http://localhost:8081](http://localhost:8081)

## Develop
These steps are only necessary if you want to help develop the app.

### Reset the database

```sh
source ./.env

./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL drop
./migrate.linux-amd64 -path db/migrations -database $POSTGRESQL_URL up
```

### Generate Dataloader code

```sh
cd backend/graph/model
go run github.com/vektah/dataloaden UserLoader string *github.com/couchless-backend/graph/model/model.User
```

### Update generated files on schema change or name change
Run in project root: `go generate ./...`\
Alternatively run this: `go run github.com/99designs/gqlgen generate`

## Changelog

See [CHANGELOG](CHANGELOG.md)

## License

[AGPL V3](LICENSE)
