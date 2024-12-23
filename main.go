package main

import (
	"database/sql"
	"github.com/kevinarchambeau/blogAggregator/internal/config"
	"github.com/kevinarchambeau/blogAggregator/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("No arguments provided")
	}

	appState := state{}
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal("Can't load config")
	}
	appState.cfg = &appConfig
	dbURL := appState.cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	appState.db = dbQueries

	cmnds := commands{
		cmds: map[string]func(*state, command) error{},
	}
	cmnds.register("login", handlerLogin)
	cmnds.register("register", handlerRegister)
	cmnds.register("reset", handlerReset)
	cmnds.register("users", handlerGetUsers)
	cmnds.register("agg", handlerAgg)
	cmnds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmnds.register("feeds", handlerGetFeeds)
	cmnds.register("follow", middlewareLoggedIn(handlerFollow))
	cmnds.register("following", middlewareLoggedIn(handlerFollowing))
	cmnds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmnds.register("browse", middlewareLoggedIn(handlerBrowse))

	cliCmd := command{}
	cliCmd.name = args[0]
	cliCmd.args = args[1:]

	err = cmnds.run(&appState, cliCmd)
	if err != nil {
		log.Printf("error running command '%s': %s\n", cliCmd.name, err)
		log.Fatal("exiting")
	}

	//fmt.Printf("app config:\n db: %s\n user: %s\n", appConfig.DbURL, appConfig.CurrentUserName)
}
