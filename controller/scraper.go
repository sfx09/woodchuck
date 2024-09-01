package controller

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/sfx09/woodchuck/internal/database"
	"github.com/sfx09/woodchuck/internal/scraper"
)

func (c *Controller) ScrapePeriodic() {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan bool)
	for {
		select {
		case <-ticker.C:
			c.ScrapeFeeds()
		case <-quit:
			return
		}
	}
}

func (c *Controller) ScrapeFeeds() {
	var wg sync.WaitGroup
	feeds, err := c.DB.GetFeedsToFetch(context.TODO(), 1)
	if err != nil {
		log.Println("Unable to fetch feeds from database")
	}
	for _, feed := range feeds {
		wg.Add(1)
		go c.ScrapeNLog(&wg, feed)
	}
	wg.Wait()
}

func (c *Controller) ScrapeNLog(wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := scraper.ScrapeFeed(feed.Url)
	if err != nil {
		log.Println("Unable to scrape url:", feed.Url)
		return
	}
	updatedFeed, err := c.DB.MarkFeedFetched(context.TODO(), database.MarkFeedFetchedParams{
		ID:          feed.ID,
		LastFetched: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		log.Println("Unable to update database:", feed.Url)
		return
	}
	log.Println("URL was scraped successfully:", updatedFeed.Url)
}
