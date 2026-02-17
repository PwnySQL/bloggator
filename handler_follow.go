package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/PwnySQL/bloggator/internal/database"
	"github.com/PwnySQL/bloggator/internal/pgerror"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects an url as single argument", cmd.name)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("%s is not a known feed URL. Please add it first using addfeed", cmd.arguments[0])
		}
		return err
	}
	follow, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			err = fmt.Errorf("You are already following feed %s (URL: %s)", feed.Name, feed.Url)
		}
		return err
	}
	fmt.Printf("Added feed '%s' for user '%s'\n", follow.FeedName, follow.UserName)

	return nil
}
