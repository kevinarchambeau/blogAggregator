# blogAggregator
Requirements:
* postgres

Create a config file `~/.gatorconfig.json` and add db config in the format of:
```json
{
  "db_url": "postgres://example"
}
```

`go build && ./blogAggregator`