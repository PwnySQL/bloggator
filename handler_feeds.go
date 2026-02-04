package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("%s does not expect arguments", cmd.name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds available")
	}

	for _, feed := range feeds {
		str := fmt.Sprintf("* %s:\n", feed.Name)
		str += fmt.Sprintf("    %s\n", feed.Url)
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		str += fmt.Sprintf("    added by: %s\n", user.Name)
		fmt.Println(str)
	}
	return nil
}
