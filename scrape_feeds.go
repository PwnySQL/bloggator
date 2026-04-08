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

func getSupportedTimeLayouts() []string {
	return []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
	}
}

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

	// Assume a default layout to avoid layout staying nil.
	// Assume that all items in a feed use the same time format.
	layout := time.RFC3339
	for _, layout_to_try := range getSupportedTimeLayouts() {
		_, err = time.Parse(layout_to_try, rssFeed.Channel.Item[0].PubDate)
		if err == nil {
			layout = layout_to_try
			break
		}
	}
	for _, rssItem := range rssFeed.Channel.Item {
		pubDate, err := time.Parse(layout, rssItem.PubDate)
		if err != nil {
			log.Printf("Could not parse pub date using layout '%s': %v", layout, err)
		}
		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       rssItem.Title,
			Url:         rssItem.Link,
			Description: sql.NullString{String: rssItem.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubDate, Valid: err == nil},
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
	// log.Println(rssFeed.String())
	return
}
