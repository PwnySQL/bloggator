package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not expect arguments", cmd.name)
	}
	err := s.db.ResetUsers(context.Background())

	return err
}
