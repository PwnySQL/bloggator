package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not expect arguments", cmd.name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	follows, err := s.db.GetFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(follows) == 0 {
		fmt.Println("No follows available")
	}

	str := fmt.Sprintf("%s follows:\n", s.cfg.CurrentUserName)
	for _, follow := range follows {
		str += fmt.Sprintf("  * %s\n", follow.FeedName)
		fmt.Println(str)
	}
	return nil
}
