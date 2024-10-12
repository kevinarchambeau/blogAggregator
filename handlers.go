package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kevinarchambeau/blogAggregator/internal/database"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no user provided, can't login")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no user provided, can't register")
	}
	currentTime := time.Now()
	result, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
	})

	err = s.cfg.SetUser(result.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' created and set as current\n", result.Name)
	fmt.Printf("Record is: %s\n", result)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.TruncateUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Users table truncated\n")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	results, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("User list: \n")
	for _, record := range results {
		if record.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", record.Name)
		} else {
			fmt.Printf("* %s\n", record.Name)
		}
	}
	fmt.Printf("\n")
	return nil
}

func handlerAgg(s *state, cmd command) error {
	//if len(cmd.args) == 0 {
	//	return fmt.Errorf("no url provided")
	//}

	data, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("RSS struct: %s\n", data)

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough args. usage: addfeed $name $url")
	}

	currentTime := time.Now()
	result, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed added: %s\n", result)

	followCmd := command{
		name: "follow",
		args: []string{cmd.args[1]},
	}
	err = handlerFollow(s, followCmd, user)
	if err != nil {
		return err
	}

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	results, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Feeds: \n")
	for _, result := range results {
		user, err := s.db.GetUserName(context.Background(), result.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s, URL: %s, user name: %s\n", result.Name, result.Url, user)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough args. usage: follow $url")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	currentTime := time.Now()
	result, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Follow created, Feed name: %s, user name: %s\n", result.FeedName, result.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	results, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%s is following: \n", s.cfg.CurrentUserName)
	for _, record := range results {
		fmt.Printf("* %s\n", record.FeedName)
	}

	return nil
}
