package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), feed.ID)

	if err != nil {
		return err
	}

	fmt.Printf("Feed follow removed. \n")
	return nil
}
