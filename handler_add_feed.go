package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/PwnySQL/bloggator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("%s expects the name and the url of the RSS Feed as argument", cmd.name)
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
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
		return err
	}
	fmt.Printf("AddedFeed:\nName: %s\nURL: %s\nUser: %s (id: %v)\n", feed.Name, feed.Url, user.Name, feed.UserID)
	return nil
}
