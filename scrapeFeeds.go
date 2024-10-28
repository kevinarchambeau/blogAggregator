package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/kevinarchambeau/blogAggregator/internal/database"
	"strings"
	"time"
)

func scrapeFeeds(s *state, duration time.Duration) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	currentSqlTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	currentTime := time.Now()

	data, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: currentSqlTime,
		ID:            feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Fetched feed: %s\n", feed.Name)
	for _, item := range data.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}
		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "posts_url_key") {
				err = s.db.UpdatePost(context.Background(), database.UpdatePostParams{
					UpdatedAt:   currentTime,
					Title:       item.Title,
					Description: item.Description,
					Url:         item.Link,
				})
			} else {
				return err
			}
		}
	}

	return nil
}
