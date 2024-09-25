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

	fmt.Printf(appConfig.DbURL)
}
