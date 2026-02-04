package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("%s does expect the RSS Feed url as argument", cmd.name)
	}
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	rssFeed.Unescape()
	rssFeed.Print()
	return nil
}
