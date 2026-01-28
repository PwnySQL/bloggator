package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not expect arguments", cmd.name)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Println("No users registered")
	}

	for _, user := range users {
		str := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			str += " (current)"
		}
		fmt.Println(str)
	}
	return nil
}
