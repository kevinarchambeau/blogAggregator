# blogAggregator
Requirements:
* postgres
* goose

Create a config file `~/.gatorconfig.json` and add db config in the format of:
```json
{
  "db_url": "postgres://postgres:$password@localhost:5432/$DB?sslmode=disable"
}
```

Note: TLS is currently disabled on purpose

* Make sure your postgres user is configured
* Create a db in postgres
* `go install github.com/pressly/goose/v3/cmd/goose@latest`
* run `goose postgres postgres://postgres:$password@localhost:5432/$DB up` from the `sql/schema` dir

To build and run: `go build && ./blogAggregator`