# blogAggregator

This is a CLI RSS aggregator that supports multiple users and feeds.

### Requirements & Setup:
* postgres
* goose


* Make sure your postgres user is configured
* Create a db in postgres
  * `CREATE DATABASE gator;` (this can be a different name as long as the connection string matches)
* `go install github.com/pressly/goose/v3/cmd/goose@latest`
* run `goose -dir sql/schema postgres postgres://postgres:$password@localhost:5432/$DB up`

Create a config file `~/.gatorconfig.json` and add db config in the format of:
```json
{
  "db_url": "postgres://postgres:$password@localhost:5432/$DB?sslmode=disable"
}
```
Note: TLS is currently disabled on purpose


To build and run: `go build && ./blogAggregator`
`go install https://github.com/kevinarchambeau/blogAggregator` should also work

### Commands
Commands will print usage info if command can't be completed

* login - sets the current user. 
  * usage: `login $user`
* register - add a user and set them as the current user. 
  * usage: `register $user`
* users - lists all users
* addfeed - add a feed to the list. 
  * usage `addfeed $name $url`
* feeds - lists all feeds
* follow - follow an existing feed another user has added. 
  * usage: `follow $url`
* following - lists all feeds the current user is following
* unfollow - remove a feed from the current user. 
  * usage: `unfollow $url`
* agg - loops through all feeds on a timer and saves posts to the db. 
  * usage: `agg $time (e.g. 60s, 1m, 1h)`
* browse - lists posts from all feeds the current user is following, sorted by date(ascending). 
  * usage: `browse $limit` (defaults to 2)
* reset - Resets the db for easy testing
