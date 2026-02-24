package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PwnySQL/bloggator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects an url as single argument", cmd.name)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("%s is not a known feed URL. Nothing to remove\n", cmd.arguments[0])
			return nil
		}
		return err
	}
	err = s.db.DeleteFeedFollow(context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Removed feed '%s' for user '%s'\n", feed.Name, user.Name)

	return nil
}
