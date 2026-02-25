package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("%s expects the time interval to fetch RSS feeds as argument", cmd.name)
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(context.Background(), s)
		if err != nil {
			return err
		}
	}
}
