# Gintama (Just Go Gin)

## Description

This repository contains the Rest API purpose. The code is written in Go.

The server is designed to be high-performant and cost-effective. The code uses dependency injection as its main design.

## Running the app locally

Create your own database name, and adjust db name on `/migrations/000001_initial-database.up.sql` file, search `dev_gintama` and replace it with your own database name.

In order to run the app, all the services have to be started.

```bash
# Run database migrations and seed test data
$ make db/migrations/up/seed
```

## Available commands

```bash
# List all commands
$ make help

# Run the app
$ make run

# Connect to the database using psql
$ make db/psql

# Create a new database migration
$ make db/migrations/new name=[name]

# Apply all up database migrations
$ make db/migrations/up

# Apply all up database migrations and seed
$ make db/migrations/up/seed

# Apply all down database migrations
$ make db/migrations/down

# Apply all refresh database migrations
$ make db/migrations/refresh

# Apply all refresh database migrations and seed
$ make db/migrations/refresh/seed

# Build application
$ make build/api
```

## Deployment

You can setup GitHub Actions refer to `github/workflows` and push the registry to your container registry, like gcr, aws ecr, etc.

### Manually

```bash
# Build application
$ make build/api
```

The command will generate two binaries:

- **Local environment binary**

  This binary is compiled for the host machine's achitecture. The binary can be found at `/bin/api`.

- **Server environment binary**

  This binary is cross-compiled for Linux with amd64 architecture. The binary can be found at `/bin/linux_amd64/api`.

You can then use the binary for deployment.
