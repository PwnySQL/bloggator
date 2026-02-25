package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/PwnySQL/bloggator/internal/database"
	"github.com/PwnySQL/bloggator/internal/pgerror"
)

func scrapeFeeds(ctx context.Context, s *state) {
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("Could not retrieve next feed to fetch: %v", err)
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
	for _, rssItem := range rssFeed.Channel.Item {
		pubDate, err := time.Parse(time.RFC3339, rssItem.PubDate)
		if err != nil {
			log.Printf("Could not parse pub date as RFC3339: %v", rssItem.PubDate, err)
		}
		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       rssFeed.Channel.Title,
			Url:         rssFeed.Channel.Link,
			Description: sql.NullString{String: rssItem.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubDate, Valid: err != nil},
			FeedID:      feed.ID,
		},
		)
		if err != nil {
			if e := pgerror.UniqueViolation(err); e != nil {
				// Trying to create the same post is not a problem.
				continue
			}
			log.Printf("Could not create post for RSS Item (%s): %v", rssItem, err)
			return
		}
	}
	return
}
