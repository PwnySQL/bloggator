package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func handlerBrowse(s *state, cmd command) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("%s expects only one optional argument", cmd.name)
	}
	numToPrint := 2
	var err error
	if len(cmd.arguments) == 1 {
		numToPrint, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("Cannot convert '%s' to number", cmd.arguments[0])
		}
	}
	posts, err := s.db.GetPosts(context.Background(), int32(numToPrint))
	if err != nil {
		return err
	}
	if len(posts) == 0 {
		fmt.Println("No follows available")
	}

	var sb strings.Builder
	for idx, post := range posts {
		sb.WriteString(fmt.Sprintf("%02d. %s", idx+1, post.Title))
		if post.PublishedAt.Valid {
			sb.WriteString(fmt.Sprintf(" (%v)", post.PublishedAt.Time.Format("2006-01-02, Mon, 03:04")))
		}
		if post.Description.Valid {
			sb.WriteString(":\n")
			sb.WriteString(fmt.Sprintf("%s\n", post.Description.String))
		}
		sb.WriteString(fmt.Sprintf("URL: %s\n", post.Url))
	}
	fmt.Println(sb.String())
	return nil
}
