package main

import (
	"context"
	"fmt"

	"github.com/PwnySQL/bloggator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not expect arguments", cmd.name)
	}
	follows, err := s.db.GetFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(follows) == 0 {
		fmt.Println("No follows available")
	}

	str := fmt.Sprintf("%s follows:\n", user.Name)
	for _, follow := range follows {
		str += fmt.Sprintf("  * %s\n", follow.FeedName)
		fmt.Println(str)
	}
	return nil
}
