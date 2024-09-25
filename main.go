package main

import (
	"fmt"
	"github.com/kevinarchambeau/blogAggregator/internal/config"
	"log"
)

func main() {
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal("Can't load config")
	}

	err = appConfig.SetUser("kevin")
	if err != nil {
		log.Fatal("Can't set user")
	}

	appConfig, err = config.Read()
	if err != nil {
		log.Fatal("Can't load config")
	}

	fmt.Printf("config:\n db: %s user: %s\n", appConfig.DbURL, appConfig.CurrentUserName)
}
