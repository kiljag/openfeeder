package main

import (
	"log"
	"openfeeder/pkg/rss"
	"time"
)

func InitFeeds() {

	feedUrls := make([]string, 0)
	feedUrls = append(feedUrls, "https://rbi.org.in/pressreleases_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/notifications_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/speeches_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/tenders_rss.xml")
	feedUrls = append(feedUrls, "https://rbi.org.in/AnnualReportMain_rss.xml")

	existsingUrls := make(map[string]bool, 0)
	feeds := rss.FetchFeedsFromDb()
	log.Printf("feeds found : %d\n", len(feeds))
	for _, feed := range feeds {
		existsingUrls[feed.Url] = true
	}

	newFeeds := make([]rss.RssFeed, 0)
	for _, url := range feedUrls {
		if existsingUrls[url] {
			continue
		}
		feed, items := rss.ParseFeedItemsFromXml(url)
		log.Println("title : ", feed.FeedHash)
		log.Println("description : ", feed.Description)
		log.Println("url : ", feed.Url)
		log.Println("feed hash : ", feed.FeedHash)
		log.Println("#items : ", len(items))
		newFeeds = append(newFeeds, feed)
	}
	log.Printf("writing %d new feeds to db\n\n", len(newFeeds))
	rss.PersistFeedsToDb(newFeeds)
}

func ProcessFeeds() {
	feeds := rss.FetchFeedsFromDb()
	log.Printf("Processing %d feeds\n", len(feeds))
	for _, feed := range feeds {
		existingItems := rss.FetchFeedItemsFromDb(feed.Id)
		existingItemHashes := make(map[string]bool)
		for _, item := range existingItems {
			existingItemHashes[item.ItemHash] = true
		}

		_, items := rss.ParseFeedItemsFromXml(feed.Url)
		newItems := make([]rss.RssFeedItem, 0)
		for _, item := range items {
			if existingItemHashes[item.ItemHash] {
				continue
			}
			item.FeedId = feed.Id
			newItems = append(newItems, item)
		}
		log.Printf("writing %d new items to db\n", len(newItems))
		rss.PersistFeedItemsToDb(newItems)
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
