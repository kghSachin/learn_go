package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kghsachin/learn_go/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenFetches time.Duration,
) {
	log.Printf("Starting scrapping with %d workers %v ", concurrency, timeBetweenFetches)
	ticker := time.NewTicker(timeBetweenFetches)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Printf("Title: %s on feed %v", item.Title, feed.Name)
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		t, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Printf("Error parsing time: %v", err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Description: description,
			Title:       item.Title,
			PublishedAt: t,
			FeedID:      feed.ID,
			Url:         item.Link,
		})
		if err != nil {
			if strings.Contains(error.Error(err), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Error creating post: %v", err)

		}
	}

	log.Printf("Fetched feed: %s", feed.Url)
	log.Printf("Fetched %d items", len(rssFeed.Channel.Item))

}
