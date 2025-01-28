package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	fmt.Printf("Current user: %s\n", user.Name)
	fmt.Println("Followed feeds:")
	for _, ff := range feedFollows {
		fmt.Printf(" * %s\n", ff.FeedName)
	}
	return nil
}
