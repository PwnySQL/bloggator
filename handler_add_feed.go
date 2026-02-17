package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/PwnySQL/bloggator/internal/database"
	"github.com/PwnySQL/bloggator/internal/pgerror"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("%s expects the name and the url of the RSS Feed as argument", cmd.name)
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			err = fmt.Errorf("Feed with URL %s is already added", url)
		}
		return err
	}
	fmt.Printf("AddedFeed:\nName: %s\nURL: %s\nUser: %s (id: %v)\n", feed.Name, feed.Url, user.Name, feed.UserID)
	_, err = s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FeedID:    feed.ID,
			UserID:    user.ID,
		},
	)
	fmt.Printf("You are now following %s\n", feed.Name)
	return nil
}
