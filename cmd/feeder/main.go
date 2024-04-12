package main

import (
	"log"
	"openfeeder/pkg/db"
	"openfeeder/pkg/rss"
	"time"
)

func InitFeeds() {

	existsingUrls := make(map[string]bool, 0)
	feeds := db.SelectFeeds()
	log.Printf("feeds found : %d\n", len(feeds))
	for _, feed := range feeds {
		log.Println("Found url : ", feed.ID, feed.Url)
		existsingUrls[feed.Url] = true
	}
	log.Println("# Existing urls : ", len(existsingUrls))

	feedUrls := make([]string, 0)
	feedUrls = append(feedUrls, "https://rbi.org.in/pressreleases_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/notifications_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/speeches_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/tenders_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/AnnualReportMain_rss.xml")
	log.Println("# Feed urls : ", len(feedUrls))

	newFeeds := make([]*rss.RssFeed, 0)
	for _, url := range feedUrls {
		if existsingUrls[url] {
			continue
		}
		feed, items := rss.ParseFeedItemsFromXml(url)
		if feed == nil || items == nil {
			continue
		}
		newFeeds = append(newFeeds, feed)
	}
	log.Println("New feeds : ", len(newFeeds))

	// adding feeds to table
	if len(newFeeds) > 0 {
		dbFeeds := make([]db.Feed, 0)
		for _, feed := range newFeeds {
			dbFeeds = append(dbFeeds, db.Feed{
				Title:       feed.Title,
				Url:         feed.Url,
				Description: feed.Description,
				FeedHash:    feed.FeedHash,
			})
		}
		db.InsertFeeds(dbFeeds)
		for _, feed := range dbFeeds {
			log.Println("Inserted feed :", feed.ID, feed.Url)
		}
	}
}

func ProcessFeed(feed db.Feed) {
	log.Println("Processing Feed : ", feed.ID, feed.Url)
	existingItems := db.SelectFeedItemsByPublishedAtDesc(feed.ID, 100)
	existingItemsMap := make(map[string]bool, 0)
	for _, item := range existingItems {
		existingItemsMap[item.ItemHash] = true
	}
	log.Println("Existing items : ", len(existingItems))

	_, items := rss.ParseFeedItemsFromXml(feed.Url)
	newDbItems := make([]db.FeedItem, 0)
	for _, item := range items {
		if existingItemsMap[item.ItemHash] {
			continue
		}
		newDbItems = append(newDbItems, db.FeedItem{
			FeedId:      feed.ID,
			Title:       item.Title,
			Link:        item.Link,
			Description: &item.Description,
			PubDate:     item.PubDate,
			ItemHash:    item.ItemHash,
		})
	}
	log.Println("New Feed items found  : ", len(newDbItems))
	if (len(newDbItems)) > 0 {
		newDbItems = db.InsertFeedItems(newDbItems)
		log.Println("Inserted items : ", len(newDbItems))
	}
}
func ProcessFeeds() {
	feeds := db.SelectFeeds()
	log.Printf("Processing %d feeds\n", len(feeds))
	for _, feed := range feeds {
		ProcessFeed(feed)
	}
	log.Printf("Processing done\n\n")
}

func main() {
	InitFeeds()
	for {
		ProcessFeeds()
		time.Sleep(5 * time.Minute)
	}
}
