package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevinarchambeau/blogAggregator/internal/database"
	"time"
)

func scrapeFeeds(s *state, duration time.Duration) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: currentTime,
		ID:            feed.ID,
	})
	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	fmt.Printf("Fetched feed: %s\n", feed.Name)
	for _, item := range data.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	fmt.Printf("\n")

	return nil
}
