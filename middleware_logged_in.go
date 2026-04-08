package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PwnySQL/bloggator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	loginStrategy := func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("No user registered with name '%s'. Register first", s.cfg.CurrentUserName)
			}
			return err
		}
		return handler(s, cmd, user)
	}
	return loginStrategy
}
