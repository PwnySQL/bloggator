package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/PwnySQL/bloggator/internal/database"
)

func scrapeFeeds(ctx context.Context, s *state) {
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("Could not retrieve next feed '%s' to fetch: %v", feed.Name, err)
		return
	}
	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{ID: feed.ID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	if err != nil {
		log.Printf("Could not mark feed '%s' as fetched: %v", feed.Name, err)
		return
	}
	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("Could not fetch feed '%s': %v", feed.Name, err)
		return
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
	rssFeed.Unescape()
	rssFeed.Print()
	return
}
