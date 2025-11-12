
# Setup

Below are the steps on how to setup the web application.

## Database Setup

To setup database we are going to use postgres, follow the steps in order:

1. Enter the postgres shell: 
    - Mac: `psql postgres`
    - Linux: `sudo -u postgres psql`
2. Create the database: `CREATE DATABASE accounting_journal;`
3. Try connecting to the database (optional): `\c accounting_journal`
4. Set Password (Linux only): `ALTER USER postgres with PASSWORD <your_pass>;`

## Goose Setup

Install goose using the following command:

```go
go install github.com/pressly/goose/v3/cmd/goose@latest
```

The default connection string:

```
postgres://postgres:postgres@localhost:5432/accounting_journal
```

`cd` into `sql/schema` directory and run:

```
goose postgres <connection_string> up
```

This will create all the required tables.

### Issues

If an error occured along the lines of `SSL Mode`, use the following to connect string to the database:

```
postgres://postgres:postgres@localhost:5432/accounting_journal?sslmode=disable
```
