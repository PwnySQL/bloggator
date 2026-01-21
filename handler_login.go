package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects the username as argument", cmd.name)
	}
	s.cfg.SetUser(cmd.arguments[0])
	fmt.Printf("Set user to %s\n", cmd.arguments[0])
	return nil
}
