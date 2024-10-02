package main

import (
	"fmt"
	"github.com/kevinarchambeau/blogAggregator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("No arguments provided")
	}

	if args[0] == "login" && len(args) != 2 {
		log.Fatal("No username provided")
	}

	appState := state{}
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal("Can't load config")
	}
	appState.cfg = &appConfig

	cmnds := commands{
		cmds: map[string]func(*state, command) error{},
	}
	cmnds.register("login", handlerLogin)

	cliCmd := command{}
	cliCmd.name = args[0]
	cliCmd.args = []string{args[1]}

	err = cmnds.run(&appState, cliCmd)
	if err != nil {
		log.Printf("error running command '%s': %s\n", cliCmd.name, err)
	}

	fmt.Printf("config:\n db: %s\n user: %s\n", appConfig.DbURL, appConfig.CurrentUserName)
}
