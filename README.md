Pismo Accessment Test

### How to run

Start database by running:

```bash
$ docker compose up -d
```

You can run the application on docker compose:

```bash
$ docker compose --profile app up -d
```

Optionally you can run it locally:

```bash
$ cd src
$ go run main.go
```

If you use VSCode, you can use the debug mode located at `/.vscode/launch.json`

### Features

- Create new accounts
- Get account by id
- Create account transactions
- List account transactions

All endpoints are located in the Postman Collection at `postman_collection.json`

### Src Structure

```
|-src
    |-main.go - Entrypoint
    |-common - Core components, utilities
    |-server - Middleware and routes
    |-account - Account model and endpoint handlers
    |-transaction - Transaction model and endpoint handlers
```

### Technologies

- Go
- Mysql
- Gorilla Mux
- Gorm