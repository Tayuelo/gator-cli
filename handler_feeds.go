package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error retrieving user with ID: %s", feed.UserID)
		}
		println("------------------------------------------------")
		printFeed(feed, user)
	}

	return nil
}
