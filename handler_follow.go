package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	printFeedFollow(feedFollow)
	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf(" * Feed:    %v\n", feedFollow.FeedName)
	fmt.Printf(" * User:    %v\n", feedFollow.UserName)
}
