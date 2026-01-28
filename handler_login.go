package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects the username as argument", cmd.name)
	}
	name := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User '%s' is not registered. Register first before logging in!", name)
	}
	s.cfg.SetUser(name)
	fmt.Printf("Set user to %s\n", cmd.arguments[0])
	return nil
}
